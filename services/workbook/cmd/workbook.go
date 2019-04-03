package main

import (
	"fmt"
	"net"
	"os"
	"syscall"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
	"google.golang.org/grpc"

	"github.com/dimdiden/portanizer-micro/services/users"
	usersgrpc "github.com/dimdiden/portanizer-micro/services/users/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	workbooksvc "github.com/dimdiden/portanizer-micro/services/workbook/implementation"
	"github.com/dimdiden/portanizer-micro/services/workbook/pb"
	"github.com/dimdiden/portanizer-micro/services/workbook/postgresdb"
	"github.com/dimdiden/portanizer-micro/services/workbook/transport"
	wbgrpc "github.com/dimdiden/portanizer-micro/services/workbook/transport/grpc"

	"os/signal"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	_ "github.com/lib/pq"
)

type config struct {
	DBConn        string `envconfig:"DB_CONNECTION"`
	GRPCAddr      string `envconfig:"GRPC_ADDR"`
	UsersGRPCAddr string `envconfig:"USERS_GRPC_ADDR"`
	Secret        string `envconfig:"SECRET"`
}

func main() {
	var cfg config
	envconfig.MustProcess("", &cfg)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "workbook",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var repository workbook.Repository
	{
		var err error
		repository, err = postgresdb.NewRepository(cfg.DBConn, logger)
		if err != nil {
			logger.Log("postgresdb", "connection failed", "err", err)
			os.Exit(1)
		}
	}
	level.Info(logger).Log("msg", "database connected")

	var usersservice users.Service
	{
		usersconn, err := grpc.Dial(cfg.UsersGRPCAddr, grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		usersservice = usersgrpc.NewGRPCClient(usersconn, logger)
	}

	var wbservice workbook.Service
	{
		wbservice = workbooksvc.NewService(usersservice, repository, logger)
	}

	var grpcServer pb.WorkbookServer
	{
		kf := func(token *jwt.Token) (interface{}, error) { return []byte(cfg.Secret), nil }
		endpoints := transport.MakeEndpoints(wbservice)
		grpcServer = wbgrpc.NewGRPCServer(kf, endpoints, logger)
	}

	var g run.Group
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", cfg.GRPCAddr)

			baseServer := grpc.NewServer()
			pb.RegisterWorkbookServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}
