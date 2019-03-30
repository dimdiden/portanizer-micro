package transport

import (
	"github.com/dimdiden/portanizer-micro/services/workbook"
	"github.com/go-kit/kit/endpoint"
)

// compile time assertions for our response types implementing endpoint.Failer.
var (
	_ endpoint.Failer = CreatePostResponse{}
)

// CreatePostRequest holds the request parameters for the CreatePost method.
type CreatePostRequest struct {
	Post workbook.Post
}

// CreatePostResponse holds the response values for the CreatePost method.
type CreatePostResponse struct {
	Post *workbook.Post
	Err  error `json:"error,omitempty"`
}

func (r CreatePostResponse) Failed() error { return r.Err }
