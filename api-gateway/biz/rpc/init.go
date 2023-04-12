package rpc

import (
	"github.com/anxiu0101/openiot-hub/kitex_gen/openiot/user/openiotuserservice"
	"github.com/anxiu0101/openiot-hub/pkg/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	userClient openiotuserservice.Client
)

func Setup() {
	initUserRpc()
}

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{consts.EtcdEndpoints})

	if err != nil {
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
		panic(err)
	}

	userClient = c
}
