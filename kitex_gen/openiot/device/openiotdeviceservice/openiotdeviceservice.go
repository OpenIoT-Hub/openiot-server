// Code generated by Kitex v0.4.4. DO NOT EDIT.

package openiotdeviceservice

import (
	"context"
	"fmt"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/device"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return openiotDeviceServiceServiceInfo
}

var openiotDeviceServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "OpeniotDeviceService"
	handlerType := (*device.OpeniotDeviceService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateDevice": kitex.NewMethodInfo(createDeviceHandler, newCreateDeviceArgs, newCreateDeviceResult, false),
		"RemoveDevice": kitex.NewMethodInfo(removeDeviceHandler, newRemoveDeviceArgs, newRemoveDeviceResult, false),
		"UpdateDevice": kitex.NewMethodInfo(updateDeviceHandler, newUpdateDeviceArgs, newUpdateDeviceResult, false),
		"GetDevice":    kitex.NewMethodInfo(getDeviceHandler, newGetDeviceArgs, newGetDeviceResult, false),
		"ListDevice":   kitex.NewMethodInfo(listDeviceHandler, newListDeviceArgs, newListDeviceResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "idl",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(device.CreateDeviceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(device.OpeniotDeviceService).CreateDevice(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateDeviceArgs:
		success, err := handler.(device.OpeniotDeviceService).CreateDevice(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateDeviceResult)
		realResult.Success = success
	}
	return nil
}
func newCreateDeviceArgs() interface{} {
	return &CreateDeviceArgs{}
}

func newCreateDeviceResult() interface{} {
	return &CreateDeviceResult{}
}

type CreateDeviceArgs struct {
	Req *device.CreateDeviceReq
}

func (p *CreateDeviceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(device.CreateDeviceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateDeviceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateDeviceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateDeviceArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateDeviceArgs) Unmarshal(in []byte) error {
	msg := new(device.CreateDeviceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateDeviceArgs_Req_DEFAULT *device.CreateDeviceReq

func (p *CreateDeviceArgs) GetReq() *device.CreateDeviceReq {
	if !p.IsSetReq() {
		return CreateDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateDeviceResult struct {
	Success *device.CreateDeviceRsp
}

var CreateDeviceResult_Success_DEFAULT *device.CreateDeviceRsp

func (p *CreateDeviceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(device.CreateDeviceRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateDeviceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateDeviceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateDeviceResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateDeviceResult) Unmarshal(in []byte) error {
	msg := new(device.CreateDeviceRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateDeviceResult) GetSuccess() *device.CreateDeviceRsp {
	if !p.IsSetSuccess() {
		return CreateDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*device.CreateDeviceRsp)
}

func (p *CreateDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func removeDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(device.RemoveDeviceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(device.OpeniotDeviceService).RemoveDevice(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RemoveDeviceArgs:
		success, err := handler.(device.OpeniotDeviceService).RemoveDevice(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RemoveDeviceResult)
		realResult.Success = success
	}
	return nil
}
func newRemoveDeviceArgs() interface{} {
	return &RemoveDeviceArgs{}
}

func newRemoveDeviceResult() interface{} {
	return &RemoveDeviceResult{}
}

type RemoveDeviceArgs struct {
	Req *device.RemoveDeviceReq
}

func (p *RemoveDeviceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(device.RemoveDeviceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RemoveDeviceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RemoveDeviceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RemoveDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RemoveDeviceArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RemoveDeviceArgs) Unmarshal(in []byte) error {
	msg := new(device.RemoveDeviceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RemoveDeviceArgs_Req_DEFAULT *device.RemoveDeviceReq

func (p *RemoveDeviceArgs) GetReq() *device.RemoveDeviceReq {
	if !p.IsSetReq() {
		return RemoveDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RemoveDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type RemoveDeviceResult struct {
	Success *device.RemoveDeviceRsp
}

var RemoveDeviceResult_Success_DEFAULT *device.RemoveDeviceRsp

func (p *RemoveDeviceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(device.RemoveDeviceRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RemoveDeviceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RemoveDeviceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RemoveDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RemoveDeviceResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RemoveDeviceResult) Unmarshal(in []byte) error {
	msg := new(device.RemoveDeviceRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RemoveDeviceResult) GetSuccess() *device.RemoveDeviceRsp {
	if !p.IsSetSuccess() {
		return RemoveDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RemoveDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*device.RemoveDeviceRsp)
}

func (p *RemoveDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(device.UpdateDeviceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(device.OpeniotDeviceService).UpdateDevice(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateDeviceArgs:
		success, err := handler.(device.OpeniotDeviceService).UpdateDevice(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateDeviceResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateDeviceArgs() interface{} {
	return &UpdateDeviceArgs{}
}

func newUpdateDeviceResult() interface{} {
	return &UpdateDeviceResult{}
}

type UpdateDeviceArgs struct {
	Req *device.UpdateDeviceReq
}

func (p *UpdateDeviceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(device.UpdateDeviceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateDeviceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateDeviceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateDeviceArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateDeviceArgs) Unmarshal(in []byte) error {
	msg := new(device.UpdateDeviceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateDeviceArgs_Req_DEFAULT *device.UpdateDeviceReq

func (p *UpdateDeviceArgs) GetReq() *device.UpdateDeviceReq {
	if !p.IsSetReq() {
		return UpdateDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateDeviceResult struct {
	Success *device.UpdateDeviceRsp
}

var UpdateDeviceResult_Success_DEFAULT *device.UpdateDeviceRsp

func (p *UpdateDeviceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(device.UpdateDeviceRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateDeviceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateDeviceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateDeviceResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateDeviceResult) Unmarshal(in []byte) error {
	msg := new(device.UpdateDeviceRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateDeviceResult) GetSuccess() *device.UpdateDeviceRsp {
	if !p.IsSetSuccess() {
		return UpdateDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*device.UpdateDeviceRsp)
}

func (p *UpdateDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(device.GetDeviceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(device.OpeniotDeviceService).GetDevice(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetDeviceArgs:
		success, err := handler.(device.OpeniotDeviceService).GetDevice(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetDeviceResult)
		realResult.Success = success
	}
	return nil
}
func newGetDeviceArgs() interface{} {
	return &GetDeviceArgs{}
}

func newGetDeviceResult() interface{} {
	return &GetDeviceResult{}
}

type GetDeviceArgs struct {
	Req *device.GetDeviceReq
}

func (p *GetDeviceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(device.GetDeviceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetDeviceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetDeviceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetDeviceArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetDeviceArgs) Unmarshal(in []byte) error {
	msg := new(device.GetDeviceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetDeviceArgs_Req_DEFAULT *device.GetDeviceReq

func (p *GetDeviceArgs) GetReq() *device.GetDeviceReq {
	if !p.IsSetReq() {
		return GetDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetDeviceResult struct {
	Success *device.GetDeviceRsp
}

var GetDeviceResult_Success_DEFAULT *device.GetDeviceRsp

func (p *GetDeviceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(device.GetDeviceRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetDeviceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetDeviceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetDeviceResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetDeviceResult) Unmarshal(in []byte) error {
	msg := new(device.GetDeviceRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetDeviceResult) GetSuccess() *device.GetDeviceRsp {
	if !p.IsSetSuccess() {
		return GetDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*device.GetDeviceRsp)
}

func (p *GetDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func listDeviceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(device.UpdateDeviceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(device.OpeniotDeviceService).ListDevice(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListDeviceArgs:
		success, err := handler.(device.OpeniotDeviceService).ListDevice(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListDeviceResult)
		realResult.Success = success
	}
	return nil
}
func newListDeviceArgs() interface{} {
	return &ListDeviceArgs{}
}

func newListDeviceResult() interface{} {
	return &ListDeviceResult{}
}

type ListDeviceArgs struct {
	Req *device.UpdateDeviceReq
}

func (p *ListDeviceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(device.UpdateDeviceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListDeviceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListDeviceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListDeviceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListDeviceArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListDeviceArgs) Unmarshal(in []byte) error {
	msg := new(device.UpdateDeviceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListDeviceArgs_Req_DEFAULT *device.UpdateDeviceReq

func (p *ListDeviceArgs) GetReq() *device.UpdateDeviceReq {
	if !p.IsSetReq() {
		return ListDeviceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListDeviceArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListDeviceResult struct {
	Success *device.ListDeviceRsp
}

var ListDeviceResult_Success_DEFAULT *device.ListDeviceRsp

func (p *ListDeviceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(device.ListDeviceRsp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListDeviceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListDeviceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListDeviceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListDeviceResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListDeviceResult) Unmarshal(in []byte) error {
	msg := new(device.ListDeviceRsp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListDeviceResult) GetSuccess() *device.ListDeviceRsp {
	if !p.IsSetSuccess() {
		return ListDeviceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListDeviceResult) SetSuccess(x interface{}) {
	p.Success = x.(*device.ListDeviceRsp)
}

func (p *ListDeviceResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateDevice(ctx context.Context, Req *device.CreateDeviceReq) (r *device.CreateDeviceRsp, err error) {
	var _args CreateDeviceArgs
	_args.Req = Req
	var _result CreateDeviceResult
	if err = p.c.Call(ctx, "CreateDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RemoveDevice(ctx context.Context, Req *device.RemoveDeviceReq) (r *device.RemoveDeviceRsp, err error) {
	var _args RemoveDeviceArgs
	_args.Req = Req
	var _result RemoveDeviceResult
	if err = p.c.Call(ctx, "RemoveDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateDevice(ctx context.Context, Req *device.UpdateDeviceReq) (r *device.UpdateDeviceRsp, err error) {
	var _args UpdateDeviceArgs
	_args.Req = Req
	var _result UpdateDeviceResult
	if err = p.c.Call(ctx, "UpdateDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetDevice(ctx context.Context, Req *device.GetDeviceReq) (r *device.GetDeviceRsp, err error) {
	var _args GetDeviceArgs
	_args.Req = Req
	var _result GetDeviceResult
	if err = p.c.Call(ctx, "GetDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListDevice(ctx context.Context, Req *device.UpdateDeviceReq) (r *device.ListDeviceRsp, err error) {
	var _args ListDeviceArgs
	_args.Req = Req
	var _result ListDeviceResult
	if err = p.c.Call(ctx, "ListDevice", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
