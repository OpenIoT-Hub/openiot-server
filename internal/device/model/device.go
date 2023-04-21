package model

import (
	"github.com/OpenIoT-Hub/openiot-server/pkg/consts"
	"github.com/OpenIoT-Hub/openiot-server/pkg/errno"
	"gorm.io/gorm"

	"errors"
)

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

func GetDeviceByID(id uint64) (BasicDevice, error) {
	var (
		device BasicDevice
	)

	if err := db.Table(consts.DeviceTableName).
		Where("id = ?", id).Find(&device).Error; err != nil {
		return device, err
	}

	return device, nil
}

func RemoveDeviceByID(id uint64) error {
	var (
		count  int64
		device BasicDevice
	)

	if err := db.Table(consts.DeviceTableName).Where("id = ?", id).Count(&count); err != nil {
		return err.Error
	} else if count < 1 {
		return errors.New(errno.GetMsg(errno.ErrorDeviceNotExist))
	}

	if err := db.Table(consts.DeviceTableName).Where("id = ?", id).Delete(&device).Error; err != nil {
		return errors.New(errno.GetMsg(errno.ErrorDatabaseOperate))
	}

	return nil
}
