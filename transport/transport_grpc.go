package transport

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"
	"github.com/rosspatil/go-kit-example/endpoint"
	pb "github.com/rosspatil/go-kit-example/pb"
)

type grpcHandlers struct {
	GetByIDHandler     grpc.Handler
	RegisterHandler    grpc.Handler
	UpdateEmailHandler grpc.Handler
	DeleteHandler      grpc.Handler
}

// NewGRPC -
func NewGRPC(e endpoint.Endpoint) pb.ServiceServer {
	return grpcHandlers{
		GetByIDHandler:     grpc.NewServer(e.GetByID, decodeGetByIDRequestGRPC, encodeGetByIDResponseGRPC),
		RegisterHandler:    grpc.NewServer(e.Register, decodeRegisterRequestGRPC, encodeRegisterResponseGRPC),
		UpdateEmailHandler: grpc.NewServer(e.UpdateEmail, decodeUpdateEmailRequestGRPC, encodeErrorOnlyResponseGRPC),
		DeleteHandler:      grpc.NewServer(e.Delete, decodeDeleteRequestGRPC, encodeErrorOnlyResponseGRPC),
	}
}

func (g grpcHandlers) GetByID(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := g.GetByIDHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}

func (g grpcHandlers) UpdateEmail(ctx context.Context, r *pb.UpdateEmailRequest) (*pb.ErrorOnlyResponse, error) {
	_, resp, err := g.UpdateEmailHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ErrorOnlyResponse), nil
}
func (g grpcHandlers) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, resp, err := g.RegisterHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RegisterResponse), nil
}

func (g grpcHandlers) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.ErrorOnlyResponse, error) {
	_, resp, err := g.DeleteHandler.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ErrorOnlyResponse), nil
}

func decodeRegisterRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RegisterRequest)
	return endpoint.RegisterRequest{Employee: *req.Employee}, nil
}

func decodeUpdateEmailRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateEmailRequest)
	return endpoint.UpdateEmailRequest{ID: req.GetId(), Email: req.GetEmail()}, nil
}

func decodeGetByIDRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRequest)
	return endpoint.GetRequest{
		ID: req.GetId(),
	}, nil
}

func decodeDeleteRequestGRPC(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteRequest)
	return endpoint.DeleteRequest{ID: req.GetId()}, nil
}

func encodeGetByIDResponseGRPC(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetResponse)
	return &pb.GetResponse{
		Employee: &resp.Employee,
	}, resp.Error
}

func encodeErrorOnlyResponseGRPC(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.ErrorOnlyResponse)

	return &pb.ErrorOnlyResponse{}, resp.Error
}

func encodeRegisterResponseGRPC(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.RegisterResponse)
	return &pb.RegisterResponse{Id: resp.ID}, resp.Error
}
