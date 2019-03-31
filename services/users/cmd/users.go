package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/oklog/run"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/dimdiden/portanizer-micro/services/users"
	"github.com/dimdiden/portanizer-micro/services/users/pb"
	"github.com/dimdiden/portanizer-micro/services/users/transport"
	grpctransport "github.com/dimdiden/portanizer-micro/services/users/transport/grpc"

	// "github.com/dimdiden/portanizer-micro/services/users/gormdb"
	userssvc "github.com/dimdiden/portanizer-micro/services/users/implementation"
	"github.com/dimdiden/portanizer-micro/services/users/mongodb"
)

type config struct {
	GRPCAddr  string `envconfig:"GRPC_ADDR"`
	MongoAddr string `envconfig:"MONGO_ADDR"`
	Secret    string `envconfig:"SECRET"`
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
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		repository, err = mongodb.NewRepository(ctx, cfg.MongoAddr, logger)
		if err != nil {
			logger.Log("mongodb", "connection failed", "err", err)
			cancel()
			os.Exit(1)
		}
	}
	level.Info(logger).Log("msg", "database connected")

	var service users.Service
	{
		service = userssvc.NewService(repository, logger)
	}

	var grpcServer pb.UsersServer
	{
		kf := func(token *jwt.Token) (interface{}, error) { return []byte(cfg.Secret), nil }
		endpoints := transport.MakeEndpoints(service)
		grpcServer = grpctransport.NewGRPCServer(kf, endpoints, logger)
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
