package grpc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/auth/pb"
	"github.com/dimdiden/portanizer-micro/services/auth/transport"
)

type grpcServer struct {
	issueTokens kitgrpc.Handler
}

func NewGRPCServer(endpoints transport.Endpoints, logger log.Logger) pb.AuthServer {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
	}
	return &grpcServer{
		issueTokens: kitgrpc.NewServer(
			endpoints.IssueTokensEndpoint,
			decodeGRPCIssueTokensRequest,
			encodeGRPCIssueTokensResponse,
			options...,
		),
	}
}

func (s *grpcServer) IssueTokens(ctx context.Context, req *pb.IssueTokensRequest) (*pb.IssueTokensReply, error) {
	_, rep, err := s.issueTokens.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.IssueTokensReply), nil
}

func decodeGRPCIssueTokensRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.IssueTokensRequest)
	return transport.IssueTokensRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

// TODO: fix BUG in case user is nil!
func encodeGRPCIssueTokensResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.IssueTokensResponse)
	if resp.Tokens == nil {
		return &pb.IssueTokensReply{
			Tokens: &pb.Tokens{}, Err: err2str(resp.Err)}, nil
	}
	return &pb.IssueTokensReply{
		Tokens: &pb.Tokens{
			Uid:    resp.Tokens.UID,
			Access: resp.Tokens.Access}, Err: err2str(resp.Err)}, nil
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
