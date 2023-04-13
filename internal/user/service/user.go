package service

import (
	"github.com/OpenIoT-Hub/openiot-server/internal/user/model"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/pack"
	api "github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
)

func GetUserInfo(userId uint) (api.UserInfo, int) {
	var (
		info      api.UserInfo
		positions []string
	)

	// Check the input uid have mapped User.
	if ok := model.CheckUserByID(userId); !ok {
		return info, errno.ErrorUserNotExist
	}

	// Search pack.User from database
	user, err := model.GetUserInfoByID(userId)
	if err != errno.Success {
		return info, errno.Error
	}

	// FIXME Read positions data from cache firstly.
	// Cache return empty value.
	if len(positions) == 0 {
		positions, err = model.GetPositionsByID(userId)
		if err != errno.Success {
			return info, errno.Error
		}
	}

	// combine the pack.User and position string array to the api.User
	info = pack.BuildUserInfo(user, positions)

	return info, errno.Success
}
