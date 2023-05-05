// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_v1 := _api.Group("/v1", _v1Mw()...)
			_v1.GET("/user", append(_listuserinfoMw(), user.ListUserInfo)...)
			_v1.POST("/user", append(_createuserMw(), user.CreateUser)...)
			_user := _v1.Group("/user", _userMw()...)
			_user.DELETE("/{id}", append(_removeuserMw(), user.RemoveUser)...)
			_user.PUT("/{id}", append(_updateuserMw(), user.UpdateUser)...)
			_user.GET("/{id}", append(_getuserinfoMw(), user.GetUserInfo)...)
		}
	}
}
