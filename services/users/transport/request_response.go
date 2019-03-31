package transport

import (
	"github.com/dimdiden/portanizer-micro/services/users"
	"github.com/go-kit/kit/endpoint"
)

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = CreateAccountResponse{}
	_ endpoint.Failer = SearchByCredsResponse{}
	_ endpoint.Failer = SearchByIDResponse{}
)

type CreateAccountRequest struct {
	Email string
	Pwd   string
}

type CreateAccountResponse struct {
	User *users.User `json:"user"`
	Err  error       `json:"error,omitempty"`
}

func (r CreateAccountResponse) Failed() error { return r.Err }

type SearchByCredsRequest struct {
	Email string
	Pwd   string
}

type SearchByCredsResponse struct {
	User *users.User
	Err  error
}

func (r SearchByCredsResponse) Failed() error { return r.Err }

type SearchByIDRequest struct {
	ID string
}

type SearchByIDResponse struct {
	User *users.User
	Err  error `json:"error,omitempty"`
}

func (r SearchByIDResponse) Failed() error { return r.Err }
