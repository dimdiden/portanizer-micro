package grpc

import (
	"context"

	"google.golang.org/grpc"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	"github.com/dimdiden/portanizer-micro/services/workbook/pb"
	"github.com/dimdiden/portanizer-micro/services/workbook/transport"
)

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) workbook.Service {

	options := []kitgrpc.ClientOption{
		kitgrpc.ClientBefore(kitjwt.ContextToGRPC()),
	}

	createPostEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Workbook",
		"CreatePost",
		encodeGRPCCreatePostRequest,
		decodeGRPCCreatePostResponse,
		pb.CreatePostReply{},
		options...,
	).Endpoint()

	return transport.Endpoints{
		CreatePostEndpoint: createPostEndpoint,
	}
}

func encodeGRPCCreatePostRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.CreatePostRequest)

	var tags []*pb.Tag
	for _, tag := range req.Post.Tags {
		t := &pb.Tag{Name: tag.Name}
		tags = append(tags, t)
	}

	return &pb.CreatePostRequest{
		Post: &pb.Post{
			Title:   req.Post.Title,
			Content: req.Post.Content,
			Tags:    tags}}, nil
}

func decodeGRPCCreatePostResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreatePostReply)

	var tags []workbook.Tag
	for _, tag := range reply.Post.Tags {
		t := workbook.Tag{Name: tag.Name}
		tags = append(tags, t)
	}

	return transport.CreatePostResponse{
		Post: &workbook.Post{
			ID:      reply.Post.Id,
			Title:   reply.Post.Title,
			Content: reply.Post.Content,
			Tags:    tags}, Err: str2err(reply.Err)}, nil // BUG
}
