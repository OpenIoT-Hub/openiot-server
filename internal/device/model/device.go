package model

import "gorm.io/gorm"

type (
	// BasicDevice 设备，核心类，external 存放一个 JSON，储存可变参数
	BasicDevice struct {
		gorm.Model
		Longitude float64 `gorm:"column:longitude"`
		Latitude  float64 `gorm:"column:latitude"`
		Version   int64   `gorm:"column:version"`  // 可变参数版本
		State     int64   `gorm:"column:state"`    // 设备状态码
		Comment   string  `gorm:"column:comment"`  // 备注
		External  string  `gorm:"column:external"` // 可变参数
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
