package main

import (
	"fmt"
	"github.com/anxiu0101/openiot-hub/pkg/util/jwt"
)

func main() {
	tokenStr, err := jwt.CreateToken(10)
	fmt.Printf("%v %v\n", tokenStr, err)
	_, err = jwt.ParseToken(tokenStr)
	if err != nil {
		println(err)
	}
}
