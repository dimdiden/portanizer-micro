package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/dimdiden/portanizer-micro/services/users"
	userstransport "github.com/dimdiden/portanizer-micro/services/users/transport"
	usersgrpc "github.com/dimdiden/portanizer-micro/services/users/transport/grpc"
)

type config struct {
	HTTPAddr      string `envconfig:"HTTP_ADDR"`
	UsersGRPCAddr string `envconfig:"USERS_GRPC_ADDR"`
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
			"svc", "gateway",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var h http.Handler
	{
		conn, err := grpc.Dial(cfg.UsersGRPCAddr, grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		service := usersgrpc.NewGRPCClient(conn, logger)
		level.Info(logger).Log("msg", "connected to Users GRPC server")
		usersEndpoints := userstransport.MakeEndpoints(service)
		h = NewServer(usersEndpoints, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", cfg.HTTPAddr)
		server := &http.Server{
			Addr:    cfg.HTTPAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)

}

// NewServer wires Go kit endpoints to the HTTP transport.
func NewServer(usersEndpoints userstransport.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		usersEndpoints.CreateAccountEndpoint,
		decodeCreateAccountRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return userstransport.CreateAccountRequest{Email: user.Email, Pwd: user.Password}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(endpoint.Failer); ok && e.Failed() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err.Error() {
	case users.ErrNotFound.Error():
		return http.StatusNotFound
	case users.ErrExists.Error():
		return http.StatusConflict
	case users.ErrNotValid.Error():
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
