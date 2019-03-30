package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/auth"
	"github.com/dimdiden/portanizer-micro/services/auth/pb"
	"github.com/dimdiden/portanizer-micro/services/auth/transport"
)

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) auth.Service {

	issueTokensEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Auth",
		"IssueTokens",
		encodeGRPCIssueTokensRequest,
		decodeGRPCIssueTokensResponse,
		pb.IssueTokensReply{},
	).Endpoint()

	return transport.Endpoints{
		IssueTokensEndpoint: issueTokensEndpoint,
	}
}

func encodeGRPCIssueTokensRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.IssueTokensRequest)
	return &pb.IssueTokensRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

func decodeGRPCIssueTokensResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.IssueTokensReply)
	return transport.IssueTokensResponse{
		Tokens: &auth.Tokens{
			UID:    reply.Tokens.Uid,
			Access: reply.Tokens.Access}, Err: str2err(reply.Err)}, nil // BUG
}
