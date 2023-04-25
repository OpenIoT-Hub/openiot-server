package pack

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/common"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
)

func BuildBaseRsp(err int) *common.BaseRsp {
	return &common.BaseRsp{
		StatusCode: int64(err),
		StatusMsg:  errno.GetMsg(err),
	}
}
