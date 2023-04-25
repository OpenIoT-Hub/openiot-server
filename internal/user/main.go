package main

import (
	"flag"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/golang/glog"

	"github.com/OpenIoT-Hub/openiot-server/config"
	"github.com/OpenIoT-Hub/openiot-server/internal/user/model"
	user "github.com/OpenIoT-Hub/openiot-server/kitex_gen/openiot/user/openiotuserservice"
	"github.com/OpenIoT-Hub/openiot-server/pkg/consts"
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
