package grpc

import (
	"context"
	"protos"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	protos.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

// GetUserInfo 获取用户信息
func (s *UserServer) GetUsersInfo(ctx context.Context, req *protos.User) (*protos.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		g.Log().Error(ctx, "metadata.FromIncomingContext failed")
		//return nil, status.Error(codes.InvalidArgument, "userId is required")
		md = metadata.New(map[string]string{})
	}
	userId, ok := md["user_id"]
	if !ok {
		g.Log().Error(ctx, "userId not found in metadata")
		return nil, status.Error(codes.InvalidArgument, "userId is required")
	}
	g.Log().Infof(ctx, "userId: %v", userId)
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required")
	}
	req.Name = "User_" + strconv.FormatInt(req.Id, 10)
	return req, nil
}
