package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/users"
	"github.com/dimdiden/portanizer-micro/services/users/pb"
	"github.com/dimdiden/portanizer-micro/services/users/transport"
)

type grpcServer struct {
	createAccount kitgrpc.Handler
}

func NewGRPCServer(endpoints transport.Endpoints, logger log.Logger) pb.UsersServer {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
	}
	return &grpcServer{
		createAccount: kitgrpc.NewServer(
			endpoints.CreateAccountEndpoint,
			decodeGRPCCreateAccountRequest,
			encodeGRPCCreateAccountResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	_, rep, err := s.createAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateAccountReply), nil
}

func decodeGRPCCreateAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateAccountRequest)
	return transport.CreateAccountRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

// TODO: fix BUG in case user is nil!
func encodeGRPCCreateAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.CreateAccountResponse)
	// fix:
	if resp.User == nil {
		return &pb.CreateAccountReply{
			User: nil, Err: err2str(resp.Err)}, nil
		// User: &pb.User{}, Err: err2str(resp.Err)}, nil
	}
	return &pb.CreateAccountReply{
		User: &pb.User{
			Id:    resp.User.ID,
			Email: resp.User.Email}, Err: err2str(resp.Err)}, nil
}

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) users.Service {

	createAccountEndpoint := kitgrpc.NewClient(
		conn,
		"pb.Users",
		"CreateAccount",
		encodeGRPCCreateAccountRequest,
		decodeGRPCCreateAccountResponse,
		pb.CreateAccountReply{},
	).Endpoint()

	return transport.Endpoints{
		CreateAccountEndpoint: createAccountEndpoint,
	}
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

func encodeGRPCCreateAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.CreateAccountRequest)
	return &pb.CreateAccountRequest{Email: req.Email, Pwd: req.Pwd}, nil
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
