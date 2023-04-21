package main

import (
	device "github.com/OpenIoT-Hub/openiot-server/internal/device/kitex_gen/openiot/device/openiotdeviceservice"
	"log"
)

func main() {
	svr := device.NewServer(new(OpeniotDeviceServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
