package transport

import (
	"github.com/dimdiden/portanizer-micro/users"
	"github.com/go-kit/kit/endpoint"
)

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = CreateAccountResponse{}
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
