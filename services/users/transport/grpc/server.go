package grpc

import (
	"context"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/dimdiden/portanizer-micro/services/users/pb"
	"github.com/dimdiden/portanizer-micro/services/users/transport"
)

type grpcServer struct {
	createAccount kitgrpc.Handler
	searchByCreds kitgrpc.Handler
	searchByID    kitgrpc.Handler
}

func NewGRPCServer(kf jwt.Keyfunc, endpoints transport.Endpoints, logger log.Logger) pb.UsersServer {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorLogger(logger),
		kitgrpc.ServerBefore(kitjwt.GRPCToContext()),
	}
	return &grpcServer{
		createAccount: kitgrpc.NewServer(
			endpoints.CreateAccountEndpoint,
			decodeGRPCCreateAccountRequest,
			encodeGRPCCreateAccountResponse,
			options...,
		),
		searchByCreds: kitgrpc.NewServer(
			endpoints.SearchByCredsEndpoint,
			decodeGRPCSearchByCredsRequest,
			encodeGRPCSearchByCredsResponse,
			options...,
		),
		searchByID: kitgrpc.NewServer(
			kitjwt.NewParser(kf, jwt.SigningMethodHS256, kitjwt.StandardClaimsFactory)(endpoints.SearchByIDEndpoint),
			// endpoints.SearchByIDEndpoint,
			decodeGRPCSearchByIDRequest,
			encodeGRPCSearchByIDResponse,
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

func (s *grpcServer) SearchByCreds(ctx context.Context, req *pb.SearchByCredsRequest) (*pb.SearchByCredsReply, error) {
	_, rep, err := s.searchByCreds.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SearchByCredsReply), nil
}

func decodeGRPCSearchByCredsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SearchByCredsRequest)
	return transport.SearchByCredsRequest{Email: req.Email, Pwd: req.Pwd}, nil
}

func encodeGRPCSearchByCredsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.SearchByCredsResponse)
	// fix:
	if resp.User == nil {
		return &pb.SearchByCredsReply{
			User: nil, Err: err2str(resp.Err)}, nil
		// User: &pb.User{}, Err: err2str(resp.Err)}, nil
	}
	return &pb.SearchByCredsReply{
		User: &pb.User{
			Id:    resp.User.ID,
			Email: resp.User.Email}, Err: err2str(resp.Err)}, nil
}

func (s *grpcServer) SearchByID(ctx context.Context, req *pb.SearchByIDRequest) (*pb.SearchByIDReply, error) {
	_, rep, err := s.searchByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SearchByIDReply), nil
}

func decodeGRPCSearchByIDRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SearchByIDRequest)
	return transport.SearchByIDRequest{ID: req.Id}, nil
}

func encodeGRPCSearchByIDResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(transport.SearchByIDResponse)

	if resp.User == nil {
		return &pb.SearchByIDReply{
			User: nil, Err: err2str(resp.Err)}, nil
		// User: &pb.User{}, Err: err2str(resp.Err)}, nil
	}
	return &pb.SearchByIDReply{
		User: &pb.User{
			Id:    resp.User.ID,
			Email: resp.User.Email}, Err: err2str(resp.Err)}, nil

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
