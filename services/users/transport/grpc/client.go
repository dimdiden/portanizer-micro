package grpc

import (
	"context"

	"google.golang.org/grpc"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/users"
	"github.com/dimdiden/portanizer-micro/services/users/pb"
	"github.com/dimdiden/portanizer-micro/services/users/transport"
)

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) users.Service {

	options := []kitgrpc.ClientOption{
		kitgrpc.ClientBefore(kitjwt.ContextToGRPC()),
	}

	createAccountEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Users",
		"CreateAccount",
		encodeGRPCCreateAccountRequest,
		decodeGRPCCreateAccountResponse,
		pb.CreateAccountReply{},
	).Endpoint()

	searchByCredsEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Users",
		"SearchByCreds",
		encodeGRPCSearchByCredsRequest,
		decodeGRPCSearchByCredsResponse,
		pb.SearchByCredsReply{},
	).Endpoint()

	searchByIDEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Users",
		"SearchByID",
		encodeGRPCSearchByIDRequest,
		decodeGRPCSearchByIDResponse,
		pb.SearchByIDReply{},
		options...,
	).Endpoint()

	return transport.Endpoints{
		CreateAccountEndpoint: createAccountEndpoint,
		SearchByCredsEndpoint: searchByCredsEndpoint,
		SearchByIDEndpoint:    searchByIDEndpoint,
	}
}

func encodeGRPCCreateAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.CreateAccountRequest)
	return &pb.CreateAccountRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

func decodeGRPCCreateAccountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateAccountReply)
	if reply.User == nil {
		return transport.CreateAccountResponse{User: nil, Err: str2err(reply.Err)}, nil
	}
	return transport.CreateAccountResponse{
		User: &users.User{
			ID:    reply.User.Id,
			Email: reply.User.Email}, Err: str2err(reply.Err)}, nil // BUG
}

// BUG: interface {} is transport.CreateAccountRequest, not transport.SearchByCredsRequest
func encodeGRPCSearchByCredsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.SearchByCredsRequest)
	return &pb.SearchByCredsRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

func decodeGRPCSearchByCredsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SearchByCredsReply)
	if reply.User == nil {
		return transport.SearchByCredsResponse{User: nil, Err: str2err(reply.Err)}, nil
	}
	return transport.SearchByCredsResponse{
		User: &users.User{
			ID:    reply.User.Id,
			Email: reply.User.Email}, Err: str2err(reply.Err)}, nil // BUG
}

func encodeGRPCSearchByIDRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.SearchByIDRequest)
	return &pb.SearchByIDRequest{Id: req.ID}, nil
}

func decodeGRPCSearchByIDResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SearchByIDReply)
	if reply.User == nil {
		return transport.SearchByIDResponse{User: nil, Err: str2err(reply.Err)}, nil
	}
	return transport.SearchByIDResponse{
		User: &users.User{
			ID:    reply.User.Id,
			Email: reply.User.Email}, Err: str2err(reply.Err)}, nil // BUG
}
