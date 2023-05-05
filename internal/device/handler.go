package main

import (
	"context"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/device"
	device "github.com/OpenIoT-Hub/openiot-server/internal/device/kitex_gen/openiot/device"
)

// OpeniotDeviceServiceImpl implements the last service interface defined in the IDL.
type OpeniotDeviceServiceImpl struct{}

// CreateDevice implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) CreateDevice(ctx context.Context, req *device.CreateDeviceReq) (resp *device.CreateDeviceRsp, err error) {
	// TODO: Your code here...
	return
}

// RemoveDevice implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) RemoveDevice(ctx context.Context, req *device.RemoveDeviceReq) (resp *device.RemoveDeviceRsp, err error) {
	// TODO: Your code here...
	return
}

// UpdateDevice implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) UpdateDevice(ctx context.Context, req *device.UpdateDeviceReq) (resp *device.UpdateDeviceRsp, err error) {
	// TODO: Your code here...
	return
}

// GetDevice implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) GetDevice(ctx context.Context, req *device.GetDeviceReq) (resp *device.GetDeviceRsp, err error) {
	// TODO: Your code here...
	return
}

// ListDevice implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) ListDevice(ctx context.Context, req *device.UpdateDeviceReq) (resp *device.ListDeviceRsp, err error) {
	// TODO: Your code here...
	return
}

// Ping implements the OpeniotDeviceServiceImpl interface.
func (s *OpeniotDeviceServiceImpl) Ping(ctx context.Context, req *device.PingReq) (resp *device.BaseRsp, err error) {
	// TODO: Your code here...
	return
}
