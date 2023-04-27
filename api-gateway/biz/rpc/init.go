package rpc

import (
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/device/openiotdeviceservice"
	"github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user/openiotuserservice"
	"github.com/OpenIoT-Hub/openiot-server/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/golang/glog"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	userClient   openiotuserservice.Client
	deviceClient openiotdeviceservice.Client
)

func Setup() {
	initUserRpc()
	initDeviceRpc()
}

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdEndpoints})

	if err != nil {
		glog.Fatal(err)
		panic(err)
	}

	c, err := openiotuserservice.NewClient(
		consts.UserServiceName,
		client.WithMuxConnection(consts.MuxConnection),
		client.WithRPCTimeout(consts.RPCTimeout),
		client.WithConnectTimeout(consts.ConnectTimeout),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		glog.Fatal(err)
		panic(err)
	}

	userClient = c
}
func initDeviceRpc() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdEndpoints})

	if err != nil {
		glog.Fatal(err)
		panic(err)
	}

	c, err := openiotdeviceservice.NewClient(
		consts.DeviceServiceName,
		client.WithMuxConnection(consts.MuxConnection),
		client.WithRPCTimeout(consts.RPCTimeout),
		client.WithConnectTimeout(consts.ConnectTimeout),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		glog.Fatal(err)
		panic(err)
	}

	deviceClient = c
}
