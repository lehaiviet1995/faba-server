package client

import (
	"context"

	"github.com/faba/syns/server/proto"
	"google.golang.org/grpc"
)

type FabaServiceClient struct {
	client proto.FabaServiceClient
}

func NewClient(cc grpc.ClientConnInterface) (FabaServiceClient, error) {

	client := FabaServiceClient{
		client: proto.NewFabaServiceClient(cc),
	}

	return client, nil
}

func (fb *FabaServiceClient) PublishSarFile(ctx context.Context, in *proto.RegisterEmployeeRequest, opts ...grpc.CallOption) (*proto.RegisterEmployeeResponse, error) {
	return fb.client.RegisterEmployee(ctx, in, opts...)
}
