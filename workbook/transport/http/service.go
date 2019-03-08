package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/auth/jwt"
	// stdjwt "github.com/dgrijalva/jwt-go"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/dimdiden/portanizer-micro/workbook/transport"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	postEndpoints transport.PostEndpoints,
	tagEndpoints transport.TagEndpoints,
	logger log.Logger,
) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(jwt.HTTPToContext()),
	}
	r.Methods("GET").Path("/posts/{id}").Handler(kithttp.NewServer(
		postEndpoints.GetByID,
		decodeGetByIDPostRequest,
		encodeResponse,
		options...,
	))
	return r
}

type errorer interface {
	Error() error
}

func decodeGetByIDPostRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.GetByIDPostRequest{ID: id}, nil
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
