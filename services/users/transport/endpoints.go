package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/dimdiden/portanizer-micro/services/users"
)

type Endpoints struct {
	CreateAccountEndpoint endpoint.Endpoint
	SearchByCredsEndpoint endpoint.Endpoint
	SearchByIDEndpoint    endpoint.Endpoint
}

func MakeEndpoints(s users.Service) Endpoints {
	return Endpoints{
		CreateAccountEndpoint: makeCreateAccountEndpoint(s),
		SearchByCredsEndpoint: makeSearchByCredsEndpoint(s),
		SearchByIDEndpoint:    makeSearchByIDEndpoint(s),
	}
}

func (e Endpoints) CreateAccount(ctx context.Context, email, pwd string) (*users.User, error) {
	resp, err := e.CreateAccountEndpoint(ctx, CreateAccountRequest{Email: email, Pwd: pwd})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateAccountResponse)
	return response.User, response.Err
}

func (e Endpoints) SearchByCreds(ctx context.Context, email, pwd string) (*users.User, error) {
	resp, err := e.SearchByCredsEndpoint(ctx, SearchByCredsRequest{Email: email, Pwd: pwd})
	if err != nil {
		return nil, err
	}
	response := resp.(SearchByCredsResponse)
	return response.User, response.Err
}

func (e Endpoints) SearchByID(ctx context.Context, id string) (*users.User, error) {
	resp, err := e.SearchByIDEndpoint(ctx, SearchByIDRequest{ID: id})
	if err != nil {
		return nil, err
	}
	response := resp.(SearchByIDResponse)
	return response.User, response.Err
}

func makeCreateAccountEndpoint(s users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		user, err := s.CreateAccount(ctx, req.Email, req.Pwd)
		return CreateAccountResponse{User: user, Err: err}, nil
	}
}

func makeSearchByCredsEndpoint(s users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchByCredsRequest)
		user, err := s.SearchByCreds(ctx, req.Email, req.Pwd)
		return SearchByCredsResponse{User: user, Err: err}, nil
	}
}

func makeSearchByIDEndpoint(s users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchByIDRequest)
		user, err := s.SearchByID(ctx, req.ID)
		return SearchByIDResponse{User: user, Err: err}, nil
	}
}
