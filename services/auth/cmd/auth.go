package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
	"google.golang.org/grpc"

	"github.com/dimdiden/portanizer-micro/services/auth"
	authsvc "github.com/dimdiden/portanizer-micro/services/auth/implementation"
	"github.com/dimdiden/portanizer-micro/services/auth/pb"
	"github.com/dimdiden/portanizer-micro/services/auth/redisdb"
	"github.com/dimdiden/portanizer-micro/services/auth/transport"
	grpctransport "github.com/dimdiden/portanizer-micro/services/auth/transport/grpc"
	"github.com/dimdiden/portanizer-micro/services/users"
	usersgrpc "github.com/dimdiden/portanizer-micro/services/users/transport/grpc"
)

type config struct {
	GRPCAddr      string `envconfig:"GRPC_ADDR"`
	UsersGRPCAddr string `envconfig:"USERS_GRPC_ADDR"`
	RedisAddr     string `envconfig:"REDIS_ADDR"`
	Secret        string `envconfig:"SECRET"`
	Expire        int    `envconfig:"EXPIRE"`
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
			"svc", "auth",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var repository auth.Repository
	{
		var err error
		repository, err = redisdb.NewRepository(cfg.RedisAddr, logger)
		if err != nil {
			logger.Log("redisdb", "connection failed", "err", err)
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
		usersservice = usersgrpc.NewGRPCClient(usersconn, logger) // <= BUG
	}

	var authservice auth.Service
	{
		authservice = authsvc.NewService(usersservice, repository, cfg.Secret, cfg.Expire, logger)
	}

	var grpcServer pb.AuthServer
	{
		endpoints := transport.MakeEndpoints(authservice)
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
			pb.RegisterAuthServer(baseServer, grpcServer)
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
