package main

import (
	"os"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/endpoint"

	"github.com/dimdiden/portanizer-micro/users"
	userstransport "github.com/dimdiden/portanizer-micro/users/transport"
)

type config struct {
	HTTPAddr string `envconfig:"HTTP_ADDR"`
}

var (
	ErrBadRouting = errors.New("bad routing")
)

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

	// r := mux.NewRouter()
	
	var h http.Handler
	{
		usersEndpoints := transport.MakePostEndpoints(postSvc)
		h = NewService(postEndpoints, tagEndpoints, logger)
	}

}

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	usersEndpoints userstransport.Endpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("POST").Path("/users").Handler(kithttp.NewServer(
		usersEndpoints.CreateAccount,
		decodeCreateAccountRequest,
		encodeResponse,
		options...,
	))
	return r
}

type errorer interface {
	Error() error
}

func decodeCreateAccountRequestRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var user users.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.Error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.Error(), w)
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
	switch err {
	case workbook.ErrNotFound:
		return http.StatusNotFound
	// case err.(*stdjwt.ValidationError):
	case jwt.ErrTokenContextMissing,
		jwt.ErrTokenExpired,
		jwt.ErrTokenInvalid,
		jwt.ErrTokenMalformed,
		jwt.ErrTokenNotActive,
		jwt.ErrUnexpectedSigningMethod:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}