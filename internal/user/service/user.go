package service

import (
	"github.com/anxiu0101/openiot-hub/internal/user/model"
	"github.com/anxiu0101/openiot-hub/pkg/errno"
	"log"
)

func GetUserInfo(userId uint) int {
	info, err := model.GetUserInfo(userId)
	if err == errno.Success {
		return errno.Success
	}
	// FIXME pack object BO -> DTO
	log.Printf("%+v\n", info)
	return 0
}
