package config

import (
	"fmt"
	"sync"

	configPkg "github.com/HughBliss/golang_background_scheduler_example.git/pkg/config"
	"github.com/HughBliss/golang_background_scheduler_example.git/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServiceName string `yaml:"service_name" env:"SERVICE_NAME" env-default:"scheduler_service"`
	ServiceVer  string `yaml:"service_version" env:"SERVICE_VERSION" env-default:"1.0.0"`

	Redis     configPkg.RedisConfig  `yaml:"redis"`
	Log       logger.LogConfig       `yaml:"logger"`
	Telemetry logger.TelemetryConfig `yaml:"telemetry"`
}

var instance *Config
var once sync.Once

func Get() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(help)
			fmt.Println(err)
		}
	})
	return instance
}
