// Code generated by hertz generator. DO NOT EDIT.

package device

import (
	device "github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/handler/device"
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
			_v1.GET("/device", append(_listdeviceMw(), device.ListDevice)...)
			_v1.POST("/device", append(_createdeviceMw(), device.CreateDevice)...)
			_device := _v1.Group("/device", _deviceMw()...)
			_device.DELETE("/{id}", append(_removedeviceMw(), device.RemoveDevice)...)
			_device.POST("/{id}", append(_updatedeviceMw(), device.UpdateDevice)...)
			_device.GET("/{id}", append(_getdeviceMw(), device.GetDevice)...)
		}
	}
}
