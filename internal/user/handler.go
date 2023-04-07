package main

import (
	"context"
	user "github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user"
)

// OpeniotUserServiceImpl implements the last service interface defined in the IDL.
type OpeniotUserServiceImpl struct{}

// Ping implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) Ping(ctx context.Context, req *user.PingReq) (resp *user.BaseRsp, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoRsp, err error) {
	// TODO: Your code here...
	return
}
