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

	// DeviceV1 The temp Device struct
	DeviceV1 struct {
		gorm.Model
		Longitude      float64 `gorm:"column:longitude"`
		Latitude       float64 `gorm:"column:latitude"`
		Capacity       float64 `gorm:"column:capacity"`
		Temperature    float64 `gorm:"column:temperature"`
		Humidity       float64 `gorm:"column:humidity"`
		CO2            float64 `gorm:"column:CO2"`
		SoundIntensity float64 `gorm:"column:soundIntensity"`
		State          int     `gorm:"column:state"`
		Info           string  `gorm:"column:info"`
	}
)
