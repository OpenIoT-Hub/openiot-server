package pack

import (
	"github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user"
	"github.com/anxiu0101/openiot-hub/pkg/errno"
)

func BuildBaseRsp(err int) *user.BaseRsp {
	return &user.BaseRsp{
		StatusCode: int64(err),
		StatusMsg:  errno.GetMsg(err),
	}
}
