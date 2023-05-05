package pack

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
)

func BuildBaseRsp(err int) *user.BaseRsp {
	return &user.BaseRsp{
		StatusCode: int64(err),
		StatusMsg:  errno.GetMsg(err),
	}
}
