package main

import (
	"fmt"
	"github.com/OpenIoT-Hub/openiot-server/pkg/util/jwt"
)

func main() {
	tokenStr, err := jwt.CreateToken(10)
	fmt.Printf("%v %v\n", tokenStr, err)
	_, err = jwt.ParseToken(tokenStr)
	if err != nil {
		println(err)
	}
}
