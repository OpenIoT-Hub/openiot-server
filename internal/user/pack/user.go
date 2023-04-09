package pack

import "github.com/anxiu0101/openiot-hub/internal/user/model"

type (
	// FIXME Position is a array and need to change/

	// UserInfo 用户信息 BO
	UserInfo struct {
		Id       uint64   `json:"id,omitempty"`        // 用户id
		Name     string   `json:"name,omitempty"`      // 用户名称
		Email    string   `json:"email,omitempty"`     // 用户工作邮箱
		PhoneNum string   `json:"phone_num,omitempty"` // 用户手机号码
		Avatar   string   `json:"avatar,omitempty"`    // 用户头像地址
		Position []string `json:"position,omitempty"`  // 用户职位，从 Authority 表中读取
	}
)

func BuildUserInfo(user model.User, position []string) UserInfo {
	return UserInfo{
		Id:       uint64(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		PhoneNum: user.PhoneNum,
		Avatar:   user.Avatar,
		Position: position,
	}
}
