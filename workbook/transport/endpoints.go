package transport

import (
	"context"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/go-kit/kit/endpoint"
)

// PostEndpoints holds all Go kit endpoints for the Post service.
type PostEndpoints struct {
	// Create  endpoint.Endpoint
	// Update  endpoint.Endpoint
	GetByID endpoint.Endpoint
	GetAll  endpoint.Endpoint
	// Delete  endpoint.Endpoint
}

// var kf = func(token *jwt.Token) (interface{}, error) { return []byte("ACCESS_SECRET_KEY"), nil }

// MakePostEndpoints initializes all Go kit endpoints for the Post service.
func MakePostEndpoints(s workbook.PostService) PostEndpoints {

	return PostEndpoints{
		// Create:  kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeCreatePostEndpoint(s)),
		// Update:  kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeUpdatePostEndpoint(s)),
		// GetByID: kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeGetByIDPostEndpoint(s)),
		GetByID: makeGetByIDPostEndpoint(s),
		// GetAll:  kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeGetAllPostEndpoint(s)),
		GetAll: makeGetAllPostEndpoint(s),
		// Delete:  kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeDeletePostEndpoint(s)),
	}
}

// func makeCreatePostEndpoint(s workbook.PostService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(CreatePostRequest)
// 		id, err := s.Create(ctx, req.Post)
// 		return CreatePostResponse{ID: id, Err: err}, nil
// 	}
// }

// func makeUpdatePostEndpoint(s workbook.PostService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(UpdatePostRequest)
// 		err := s.Update(ctx, req.ID, req.Post)
// 		return UpdatePostResponse{Err: err}, nil
// 	}
// }

func makeGetByIDPostEndpoint(s workbook.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDPostRequest)
		post, err := s.GetByID(ctx, req.ID)
		return GetByIDPostResponse{Post: post, Err: err}, nil
	}
}

func makeGetAllPostEndpoint(s workbook.PostService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(GetAllPostRequest)
		posts, err := s.GetAll(ctx)
		return GetAllPostResponse{Posts: posts, Err: err}, nil
	}
}

// func makeDeletePostEndpoint(s workbook.PostService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(DeletePostRequest)
// 		err := s.Delete(ctx, req.ID)
// 		return DeletePostResponse{Err: err}, nil
// 	}
// }

// TagEndpoints holds all Go kit endpoints for the Tag service.
type TagEndpoints struct {
	// Create endpoint.Endpoint
	// Update endpoint.Endpoint
	// GetAll endpoint.Endpoint
	// Delete endpoint.Endpoint
}

// MakeTagEndpoints initializes all Go kit endpoints for the Tag service.
func MakeTagEndpoints(s workbook.TagService) TagEndpoints {
	return TagEndpoints{
		// Create: kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeCreateTagEndpoint(s)),
		// Update: kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeUpdateTagEndpoint(s)),
		// GetAll: kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeGetAllTagEndpoint(s)),
		// Delete: kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(makeDeleteTagEndpoint(s)),
	}
}

// func makeCreateTagEndpoint(s workbook.TagService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(CreateTagRequest)
// 		id, err := s.Create(ctx, req.PostID, req.Tag)
// 		return CreateTagResponse{ID: id, Err: err}, nil
// 	}
// }

// func makeUpdateTagEndpoint(s workbook.TagService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(UpdateTagRequest)
// 		err := s.Update(ctx, req.ID, req.Tag)
// 		return UpdateTagResponse{Err: err}, nil
// 	}
// }

// func makeGetAllTagEndpoint(s workbook.TagService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		// req := request.(GetAllTagRequest)
// 		tags, err := s.GetAll(ctx)
// 		return GetAllTagResponse{Tags: tags, Err: err}, nil
// 	}
// }

// func makeDeleteTagEndpoint(s workbook.TagService) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(DeleteTagRequest)
// 		err := s.Delete(ctx, req.ID)
// 		return DeleteTagResponse{Err: err}, nil
// 	}
// }
