package transport

import (
	"context"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the Workbook service.
type Endpoints struct {
	CreatePostEndpoint endpoint.Endpoint
}

func MakeEndpoints(s workbook.Service) Endpoints {

	return Endpoints{
		CreatePostEndpoint: makeCreatePostEndpoint(s),
	}
}

func makeCreatePostEndpoint(s workbook.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePostRequest)
		post, err := s.CreatePost(ctx, req.Post)
		return CreatePostResponse{Post: post, Err: err}, nil
	}
}

func (e Endpoints) CreatePost(ctx context.Context, p workbook.Post) (*workbook.Post, error) {
	resp, err := e.CreatePostEndpoint(ctx, CreatePostRequest{Post: p})
	if err != nil {
		return nil, err
	}
	response := resp.(CreatePostResponse)
	return response.Post, response.Err
}
