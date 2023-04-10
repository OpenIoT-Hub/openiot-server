package service

import (
	api "github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user"
	"github.com/anxiu0101/openiot-hub/internal/user/model"
	"github.com/anxiu0101/openiot-hub/internal/user/pack"
	"github.com/anxiu0101/openiot-hub/pkg/errno"
)

func GetUserInfo(userId uint) (api.User, int) {
	var (
		info      api.User
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
