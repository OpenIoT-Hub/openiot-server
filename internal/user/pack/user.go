package pack

import (
	"github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user"
)

type (
	// FIXME Position is a array and need to change/

	// UserInfo 用户信息 BO
	UserInfo struct {
		Id       uint64 `json:"id,omitempty"`        // 用户id
		Name     string `json:"name,omitempty"`      // 用户名称
		Email    string `json:"email,omitempty"`     // 用户工作邮箱
		PhoneNum string `json:"phone_num,omitempty"` // 用户手机号码
		Avatar   string `json:"avatar,omitempty"`    // 用户头像地址
	}
)

func BuildUserInfo(info UserInfo, position []string) user.User {
	return user.User{
		Id:       info.Id,
		Name:     info.Name,
		Email:    info.Email,
		PhoneNum: info.PhoneNum,
		Avatar:   info.Avatar,
		// Position: position,
	}
}
