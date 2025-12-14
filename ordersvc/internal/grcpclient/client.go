package grcpclient

import (
	grpc "base_util/grpcclient"
	"protos"
	"sync"
)

var (
	UserClient     protos.UserServiceClient
	userClientOnce sync.Once

	SkuClient     protos.SkuServiceClient
	skuClientOnce sync.Once
)

func GetUserClient() protos.UserServiceClient {
	userClientOnce.Do(func() {
		UserClient = protos.NewUserServiceClient(grpc.NewGrpcClient(":9091", "usersvc"))
	})
	return UserClient
}

func GetSkuClient() protos.SkuServiceClient {
	skuClientOnce.Do(func() {
		SkuClient = protos.NewSkuServiceClient(grpc.NewGrpcClient(":9092", "skusvc"))
	})
	return SkuClient
}
