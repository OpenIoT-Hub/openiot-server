package main

import (
	"context"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/pack"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/service"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
)

// OpeniotUserServiceImpl implements the last service interface defined in the IDL.
type OpeniotUserServiceImpl struct{}

// Ping implements the OpeniotUserServiceImpl interface.
func (s *OpeniotUserServiceImpl) Ping(ctx context.Context, req *user.PingReq) (resp *user.BaseRsp, err error) {
	resp = new(user.BaseRsp)
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
