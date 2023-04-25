package main

import (
	"context"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/pack"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/service"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/common"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
)

// OpeniotUserServiceImpl implements the last service interface defined in the IDL.
type OpeniotUserServiceImpl struct{}

// Ping implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) Ping(ctx context.Context, req *common.PingReq) (resp *common.BaseRsp, err error) {
	resp = new(common.BaseRsp)
	resp.StatusCode = 200
	resp.StatusMsg = "Pong!"
	return
}

// GetUserInfo implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (resp *user.GetUserInfoRsp, err error) {
	resp = new(user.GetUserInfoRsp)
	id := req.UserId
	info, eCode := service.GetUserInfo(uint(id))
	base := pack.BuildBaseRsp(eCode)
	resp.UserInfo = &info
	resp.Base = base
	return
}

// CreateUser implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserRsp, err error) {
	// TODO: Your code here...
	return
}

// RemoveUser implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) RemoveUser(ctx context.Context, req *user.RemoveUserReq) (resp *user.RemoveUserRsp, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserRsp, err error) {
	// TODO: Your code here...
	return
}

// ListUserInfo implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) ListUserInfo(ctx context.Context, req *user.ListUserInfoReq) (resp *user.ListUserInfoRsp, err error) {
	// TODO: Your code here...
	return
}
