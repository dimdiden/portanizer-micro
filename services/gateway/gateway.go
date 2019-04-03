package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"

	authtransport "github.com/dimdiden/portanizer-micro/services/auth/transport"
	authgrpc "github.com/dimdiden/portanizer-micro/services/auth/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/users"
	userstransport "github.com/dimdiden/portanizer-micro/services/users/transport"
	usersgrpc "github.com/dimdiden/portanizer-micro/services/users/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	wbtransport "github.com/dimdiden/portanizer-micro/services/workbook/transport"
	wbgrpc "github.com/dimdiden/portanizer-micro/services/workbook/transport/grpc"
)

type config struct {
	HTTPAddr         string `envconfig:"HTTP_ADDR"`
	UsersGRPCAddr    string `envconfig:"USERS_GRPC_ADDR"`
	AuthGRPCAddr     string `envconfig:"AUTH_GRPC_ADDR"`
	WorkbookGRPCAddr string `envconfig:"WORKBOOK_GRPC_ADDR"`
	Secret           string `envconfig:"SECRET"`
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
		kf := func(token *jwt.Token) (interface{}, error) { return []byte(cfg.Secret), nil }

		usersconn, err := grpc.Dial(cfg.UsersGRPCAddr, grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		usersservice := usersgrpc.NewGRPCClient(usersconn, logger)
		level.Info(logger).Log("msg", "connected to Users GRPC server")
		usersEndpoints := userstransport.MakeEndpoints(usersservice)

		authconn, err := grpc.Dial(cfg.AuthGRPCAddr, grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		authservice := authgrpc.NewGRPCClient(authconn, logger)
		level.Info(logger).Log("msg", "connected to Auth GRPC server")
		authEndpoints := authtransport.MakeEndpoints(authservice)

		wbconn, err := grpc.Dial(cfg.WorkbookGRPCAddr, grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		wbservice := wbgrpc.NewGRPCClient(wbconn, logger)
		level.Info(logger).Log("msg", "connected to Workbook GRPC server")
		wbEndpoints := wbtransport.MakeEndpoints(wbservice)

		h = NewServer(kf, usersEndpoints, authEndpoints, wbEndpoints, logger)
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
func NewServer(
	kf jwt.Keyfunc,
	usersEndpoints userstransport.Endpoints,
	authEndpoints authtransport.Endpoints,
	wbEndpoints wbtransport.Endpoints,
	logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(kitjwt.HTTPToContext()),
	}
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		usersEndpoints.CreateAccountEndpoint,
		decodeCreateAccountRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/user").Handler(kithttp.NewServer(
		kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(usersEndpoints.SearchByIDEndpoint),
		decodeSearchByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/signin").Handler(kithttp.NewServer(
		authEndpoints.IssueTokensEndpoint,
		decodeIssueTokensRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/posts").Handler(kithttp.NewServer(
		kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(wbEndpoints.CreatePostEndpoint),
		decodeCreatePostRequest,
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

func decodeSearchByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return userstransport.SearchByIDRequest{}, nil
}

func decodeIssueTokensRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return authtransport.IssueTokensRequest{Email: user.Email, Pwd: user.Password}, nil
}

func decodeCreatePostRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var tmp struct {
		ID       string
		Title    string
		Content  string
		TagNames []string `json:"tags"`
	}
	if err := json.NewDecoder(r.Body).Decode(&tmp); err != nil {
		return nil, err
	}
	var tags []workbook.Tag
	for _, tn := range tmp.TagNames {
		tags = append(tags, workbook.Tag{Name: tn})
	}
	return wbtransport.CreatePostRequest{
		Post: workbook.Post{
			ID:      tmp.ID,
			Title:   tmp.Title,
			Content: tmp.Content,
			Tags:    tags,
		}}, nil
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
	case jwt.ErrSignatureInvalid.Error(),
		kitjwt.ErrTokenContextMissing.Error(),
		kitjwt.ErrTokenExpired.Error(),
		kitjwt.ErrTokenInvalid.Error(),
		kitjwt.ErrTokenMalformed.Error(),
		kitjwt.ErrTokenNotActive.Error(),
		kitjwt.ErrUnexpectedSigningMethod.Error():
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
