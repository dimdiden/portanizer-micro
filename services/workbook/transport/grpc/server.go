package grpc

import (
	"context"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	"github.com/dimdiden/portanizer-micro/services/workbook/pb"
	"github.com/dimdiden/portanizer-micro/services/workbook/transport"
)

type grpcServer struct {
	createPost kitgrpc.Handler
}

func NewGRPCServer(kf jwt.Keyfunc, endpoints transport.Endpoints, logger log.Logger) pb.WorkbookServer {

	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
		kitgrpc.ServerBefore(kitjwt.GRPCToContext()),
	}
	return &grpcServer{
		createPost: kitgrpc.NewServer(
			kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(endpoints.CreatePostEndpoint),
			decodeGRPCCreatePostRequest,
			encodeGRPCCreatePostResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	_, rep, err := s.createPost.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreatePostReply), nil
}

func decodeGRPCCreatePostRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreatePostRequest)

	var tags []workbook.Tag
	for _, tag := range req.Post.Tags {
		t := workbook.Tag{Name: tag.Name}
		tags = append(tags, t)
	}

	return transport.CreatePostRequest{
		Post: workbook.Post{
			ID:      req.Post.Id,
			Title:   req.Post.Title,
			Content: req.Post.Content,
			Tags:    tags,
		}}, nil
}

// TODO: fix BUG in case user is nil!
func encodeGRPCCreatePostResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.CreatePostResponse)

	if resp.Post == nil {
		return &pb.CreatePostReply{
			Post: &pb.Post{}, Err: err2str(resp.Err)}, nil
	}

	var tags []*pb.Tag
	for _, tag := range resp.Post.Tags {
		t := &pb.Tag{Id: tag.ID, Name: tag.Name}
		tags = append(tags, t)
	}

	return &pb.CreatePostReply{
		Post: &pb.Post{
			Id:      resp.Post.ID,
			Title:   resp.Post.Title,
			Content: resp.Post.Content,
			Tags:    tags}, Err: err2str(resp.Err)}, nil
}

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
