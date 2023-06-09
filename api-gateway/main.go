// Code generated by hertz generator.

package main

import (
	"context"
	"flag"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/golang/glog"

	"github.com/OpenIoT-Hub/openiot-server/api-gateway/biz/rpc"
	"github.com/OpenIoT-Hub/openiot-server/pkg/consts"
)

func init() {
	rpc.Setup()
	flag.Parse()
}

func main() {
	// write log file into dir before program exists.
	defer glog.Flush()

	svr := server.Default(
		server.WithHostPorts(consts.GatewayListenAddress),
		server.WithHandleMethodNotAllowed(true),
	)

	svr.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(errno.Success, utils.H{"message": "pong"})
	})

	register(svr)
	svr.Spin()
}
