package main

import (
	"flag"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/golang/glog"

	"github.com/anxiu0101/openiot-hub/internal/user/config"
	user "github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user/openiotuserservice"
	"github.com/anxiu0101/openiot-hub/internal/user/model"
	"github.com/anxiu0101/openiot-hub/pkg/consts"
)

func init() {
	config.Setup()
	model.Setup()
}

func main() {
	var logDir string
	flag.StringVar(&logDir, "log_dir", "./log", "log file dir addr")

	defer glog.Flush()
	addr, err := net.ResolveTCPAddr("tcp", consts.UserServiceListenAddress)
	if err != nil {
		glog.Fatal(err)
	}

	svr := user.NewServer(
		new(OpeniotUserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: consts.UserServiceName,
		}),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()
	if err != nil {
		glog.Fatal(err)
	}
}
