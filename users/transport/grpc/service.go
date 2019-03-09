package grpc

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/users"
	"github.com/dimdiden/portanizer-micro/users/pb"
	"github.com/dimdiden/portanizer-micro/users/transport"
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
	fmt.Println("func (s *grpcServer) CreateAccount(ctx context.Context, ")
	return rep.(*pb.CreateAccountReply), nil
}

// func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) users.Service {

// 	createAccountEndpoint := kitgrpc.NewClient(
// 			conn,
// 			"pb.Users",
// 			"CreateAccount",
// 			encodeGRPCCreateAccountRequest,
// 			decodeGRPCCreateAccountResponse,
// 			pb.CreateAccountReply{},
// 		).Endpoint()
	
// 	return transport.Endpoints{
// 		CreateAccountEndpoint: createAccountEndpoint,
// 	}
// }

func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) transport.Endpoints {

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

func decodeGRPCCreateAccountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateAccountRequest)
	fmt.Println("func decodeGRPCCreateAccountRequest(_ context.Context ",req.Email, req.Pwd)
	return transport.CreateAccountRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

// BUG
func decodeGRPCCreateAccountResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateAccountReply)
	fmt.Println("func decodeGRPCCreateAccountResponse(_ context.  ",reply.User.Id, reply.User.Email, reply.User.Pwd)
	return transport.CreateAccountResponse{
		User: &users.User{
			ID: reply.User.Id,
			Email: reply.User.Email,
			Password: reply.User.Pwd}, Err: str2err(reply.Err)}, nil // BUG
}

// BUG 
func encodeGRPCCreateAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(transport.CreateAccountRequest)
	fmt.Println("func encodeGRPCCreateAccountRequest(_ context.Context  ",req.Email, req.Pwd)
	return &pb.CreateAccountRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

func encodeGRPCCreateAccountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.CreateAccountResponse)
	fmt.Println("func encodeGRPCCreateAccountResponse(_ context.Context,  ",resp.User.ID, resp.User.Email, resp.User.Password)
	return &pb.CreateAccountReply{
		User: &pb.User{
			Id: resp.User.ID,
			Email: resp.User.Email,
			Pwd: resp.User.Password}, Err: err2str(resp.Err)}, nil
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

