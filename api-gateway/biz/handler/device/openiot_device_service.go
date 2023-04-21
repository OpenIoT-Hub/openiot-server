// Code generated by hertz generator.

package device

import (
	"context"

	device "github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/model/openiot/device"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateDevice .
// @router /api/v1/device [POST]
func CreateDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.CreateDeviceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(device.CreateDeviceRsp)

	c.JSON(consts.StatusOK, resp)
}

// RemoveDevice .
// @router api/v1/device/:id [DELETE]
func RemoveDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.RemoveDeviceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(device.RemoveDeviceRsp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateDevice .
// @router api/v1/device/:id [POST]
func UpdateDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.UpdateDeviceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(device.UpdateDeviceRsp)

	c.JSON(consts.StatusOK, resp)
}

// GetDevice .
// @router api/v1/device/:id [GET]
func GetDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.GetDeviceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(device.GetDeviceRsp)

	c.JSON(consts.StatusOK, resp)
}

// ListDevice .
// @router api/v1/device/:id [GET]
func ListDevice(ctx context.Context, c *app.RequestContext) {
	var err error
	var req device.UpdateDeviceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(device.ListDeviceRsp)

	c.JSON(consts.StatusOK, resp)
}
