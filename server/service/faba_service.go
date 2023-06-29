package service

import (
	"context"
	"fmt"

	"github.com/lehaiviet1995/faba-server/server/proto"
	"google.golang.org/grpc"
)

type FabaService struct {
	proto.UnimplementedFabaServiceServer
}

func NewSarFileService() *FabaService {
	return &FabaService{}
}

func Apply(server *grpc.Server) {
	proto.RegisterFabaServiceServer(
		server,
		NewSarFileService(),
	)
}

// RegisterEmployee register employee
func (fb *FabaService) RegisterEmployee(ctx context.Context, in *proto.RegisterEmployeeRequest) (*proto.RegisterEmployeeResponse, error) {
	fmt.Println("call create employee")
	protoRegisterEmployeeResponse := proto.RegisterEmployeeResponse{
		Message: "create employee success",
	}
	return &protoRegisterEmployeeResponse, nil
}
