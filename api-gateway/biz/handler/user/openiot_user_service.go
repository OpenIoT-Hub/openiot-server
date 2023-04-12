// Code generated by hertz generator.

package user

import (
	"context"

	user "github.com/anxiu0101/openiot-hub/api-gateway/biz/model/openiot/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetUserInfo 通过用户 ID 获取用户资料.
// @router /api/v1/user/info [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserInfoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.GetUserInfoRsp)

	c.JSON(consts.StatusOK, resp)
}
