package model

import "gorm.io/gorm"

type (
	Device struct {
		gorm.Model
		longitude float64
		latitude  float64
		version   int64  // 可变参数版本
		state     int64  // 设备状态码
		external  string // 可变参数
	}
)
