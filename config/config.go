package config

import (
	"github.com/jinzhu/configor"
)

// Global Config Variable
var Config = struct {
	AppMode string `default:debug`
	AppName string `default:"openiot-hub"`

	Server struct {
		ReadTimeout  int
		WriteTimeout int
	}

	DB struct {
		Name     string
		User     string
		Password string
		Port     uint
	}

	Redis struct {
		Password    string
		DbName      int
		Port        int
		MaxIdle     int
		MaxActive   int
		IdleTimeout int
	}

	RabbitMQ struct {
		RabbitMQ         string
		RabbitMQUser     string
		RabbitMQPassWord string
		RabbitMQHost     string
		RabbitMQPort     string
	}

	Avatar struct{}

	Wechat struct {
		AppID  string
		Secret string
	}

	SMS struct{}
}{}

func Setup() {
	configor.New(&configor.Config{
		Debug: true,
		// Verbose:           true,
		ErrorOnUnmatchedKeys: true,
	}).Load(&Config, "./config/config.toml")
}
