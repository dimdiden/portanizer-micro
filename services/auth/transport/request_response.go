package transport

import (
	"github.com/dimdiden/portanizer-micro/services/auth"
	"github.com/go-kit/kit/endpoint"
)

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = IssueTokensResponse{}
)

type IssueTokensRequest struct {
	Email string
	Pwd   string
}

type IssueTokensResponse struct {
	Tokens *auth.Tokens `json:"tokens"`
	Err    error        `json:"error,omitempty"`
}

func (r IssueTokensResponse) Failed() error { return r.Err }
