package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/dimdiden/portanizer-micro/services/auth"
)

type Endpoints struct {
	IssueTokensEndpoint endpoint.Endpoint
}

func MakeEndpoints(s auth.Service) Endpoints {
	return Endpoints{
		IssueTokensEndpoint: makeIssueTokensEndpoint(s),
	}
}

func (e Endpoints) IssueTokens(ctx context.Context, email, pwd string) (*auth.Tokens, error) {
	resp, err := e.IssueTokensEndpoint(ctx, IssueTokensRequest{Email: email, Pwd: pwd})
	if err != nil {
		return nil, err
	}
	response := resp.(IssueTokensResponse)
	return response.Tokens, response.Err
}

func makeIssueTokensEndpoint(s auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IssueTokensRequest)
		tokens, err := s.IssueTokens(ctx, req.Email, req.Pwd)
		return IssueTokensResponse{Tokens: tokens, Err: err}, nil
	}
}
