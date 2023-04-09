package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// UserRegister .
// @router /openiot/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req user.UserRegisterRequest
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	SendErrorResponse(c, errno.ParamError.WithMessage(err.Error()))
	//	return
	//}
	//
	//userId, token, err := rpc.UserRegister(context.Background(), &user.UserRegisterRequest{
	//	Username: req.Username,
	//	Password: req.Password,
	//})
	//
	//if err != nil {
	//	SendErrorResponse(c, err)
	//	return
	//}
	//
	//SendCommonResponse(c, &basic.UserRegisterResponse{
	//	StatusCode: errno.SuccessCode,
	//	StatusMsg:  errno.SuccessMsg,
	//	Token:      token,
	//	UserId:     userId,
	//})
}

// UserLogin .
// @router /openiot/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	//var err error
	//var req basic.UserLoginRequest
	//err = c.BindAndValidate(&req)
	//if err != nil {
	//	SendErrorResponse(c, errno.ParamError.WithMessage(err.Error()))
	//	return
	//}
	//
	//userId, token, err := rpc.UserLogin(context.Background(), &user.UserLoginRequest{
	//	Username: req.Username,
	//	Password: req.Password,
	//})
	//
	//if err != nil {
	//	SendErrorResponse(c, err)
	//	return
	//}
	//
	//SendCommonResponse(c, &basic.UserLoginResponse{
	//	StatusCode: errno.SuccessCode,
	//	StatusMag:  errno.SuccessMsg,
	//	UserId:     userId,
	//	Token:      token,
	//})
}

// GetUserInfo
// @router /api/v1/:uid [GET]
func GetUserInfo() {

}
