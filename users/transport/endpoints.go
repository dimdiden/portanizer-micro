package transport

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"github.com/dimdiden/portanizer-micro/users"
)

type Endpoints struct {
	CreateAccountEndpoint endpoint.Endpoint
}

func MakeEndpoints(s users.Service) Endpoints {
	return Endpoints{
		CreateAccountEndpoint: makeCreateAccountEndpoint(s),
	}
}

func (e Endpoints) CreateAccount(ctx context.Context, email, pwd string) (*users.User, error) {
	resp, err := e.CreateAccountEndpoint(ctx, CreateAccountRequest{Email: email, Pwd: pwd})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateAccountResponse)
	fmt.Println("func (e Endpoints) CreateAccount ", response.User.ID, response.User.Email)
	return response.User, response.Err
}

func makeCreateAccountEndpoint(s users.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		user, err := s.CreateAccount(ctx, req.Email, req.Pwd)
		fmt.Println("func makeCreateAccountEndpoint(s users.Service) ", user.ID, user.Email, user.Password)
		return CreateAccountResponse{User: user, Err: err}, nil
	}
}
