package handler

import (
	"github.com/anxiu0101/openiot-hub/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type ErrResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendErrorResponse(c *app.RequestContext, err error) {
	errno := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ErrResponse{
		Code: errno.ErrorCode,
		Msg:  errno.ErrorMsg,
	})
}

func SendCommonResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

// PhaseToken returns Userid
func PhaseToken(token string) (int64, error) {
	//return rpc.CheckToken(context.Background(), &user.CheckTokenRequest{
	//	Token: token,
	//})
	return 0, nil
}
