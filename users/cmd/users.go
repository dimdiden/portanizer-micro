package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"net"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"github.com/kelseyhightower/envconfig"

	"github.com/dimdiden/portanizer-micro/users"
	"github.com/dimdiden/portanizer-micro/users/pb"
	"github.com/dimdiden/portanizer-micro/users/transport"
	grpctransport "github.com/dimdiden/portanizer-micro/users/transport/grpc"
	"github.com/dimdiden/portanizer-micro/users/gormdb"
	userssvc "github.com/dimdiden/portanizer-micro/users/implementation"
)

type config struct {
	GRPCAddr string `envconfig:"GRPC_ADDR"`
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
			"svc", "users",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var repository users.Repository
	{
		repository = gormdb.New(logger)
	}
	level.Info(logger).Log("msg", "database connected")

	var service users.Service
	{
		service = userssvc.NewService(repository, logger)
	}

	var grpcServer pb.UsersServer 
	{
		endpoints := transport.MakeEndpoints(service)
		grpcServer = grpctransport.NewGRPCServer(endpoints, logger)
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
			pb.RegisterUsersServer(baseServer, grpcServer)
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
