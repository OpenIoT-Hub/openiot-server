package rpc

import (
	"context"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
)

func GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.GetUserInfoRsp, int) {
	resp, err := userClient.GetUserInfo(ctx, req)

	// TODO add errno
	if err != nil {
		return nil, errno.Error
	}

	// TODO add errno
	if resp.Base.StatusCode != errno.Success {
		return nil, errno.Error
	}

	return resp, 0
}
