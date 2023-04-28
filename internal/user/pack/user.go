package pack

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user"
)

type (
	// FIXME Position is a array and need to change/

	// UserBasic 用户信息 BO
	UserBasic struct {
		Id       uint64 `json:"id,omitempty"`        // 用户id
		Name     string `json:"name,omitempty"`      // 用户名称
		Email    string `json:"email,omitempty"`     // 用户工作邮箱
		PhoneNum string `json:"phone_num,omitempty"` // 用户手机号码
		Avatar   string `json:"avatar,omitempty"`    // 用户头像地址
	}

	UserInfo struct {
		UserBasic
		positions []string
	}
)

func BuildUserInfo(info UserBasic, position []string) user.UserInfo {
	return user.UserInfo{
		User: &user.User{
			Id:       info.Id,
			Name:     info.Name,
			Email:    info.Email,
			PhoneNum: info.PhoneNum,
			Avatar:   info.Avatar,
		},
		Position: position,
	}
}
