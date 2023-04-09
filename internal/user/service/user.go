package service

import (
	"github.com/anxiu0101/openiot-hub/internal/user/model"
	"github.com/anxiu0101/openiot-hub/internal/user/pack"
	"github.com/anxiu0101/openiot-hub/pkg/errno"
)

func GetUserInfo(userId uint) (pack.UserInfo, int) {
	var (
		info pack.UserInfo
	)

	user, err := model.GetUserInfoByID(userId)
	if err != errno.Success {
		return info, errno.Error
	}

	positions, err := model.GetPositionsByID(userId)
	if err != errno.Success {
		return info, errno.Error
	}

	info = pack.BuildUserInfo(user, positions)

	return info, errno.Success
}
