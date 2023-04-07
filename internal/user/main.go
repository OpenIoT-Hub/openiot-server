package main

import (
	user "github.com/anxiu0101/openiot-hub/internal/user/kitex_gen/openiot/user/openiotuserservice"
	"log"
)

func main() {
	svr := user.NewServer(new(OpeniotUserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
