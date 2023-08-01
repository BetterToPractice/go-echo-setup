package lib

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath = "./config.yml"

var defaultConfig = Config{
	Name: "go-echo-setup",
	Http: &HttpConfig{
		Host: "0.0.0.0",
		Port: 9999,
	},
	Log: &LogConfig{},
}

func NewConfig() Config {
	config := defaultConfig
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(errors.Wrap(err, "failed to read config"))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(errors.Wrap(err, "failed to marshal config"))
	}
	return config
}

type Config struct {
	Name string      `mapstructure:"Name"`
	Http *HttpConfig `mapstructure:"Http"`
	Log  *LogConfig  `mapstructure:"Log"`
}

type HttpConfig struct {
	Host string `mapstructure:"Host" validate:"ipv4"`
	Port int    `mapstructure:"Port" validate:"gte=1,lte=65535"`
}

type LogConfig struct {
	Level       string `mapstructure:"Level"`
	Format      string `mapstructure:"Format"`
	Directory   string `mapstructure:"Directory"`
	Development string `mapstructure:"Development"`
}

func (a *HttpConfig) ListenAddr() string {
	if err := validator.New().Struct(a); err != nil {
		return "0.0.0.0:5111"
	}
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func SetConfigPath(path string) {
	configPath = path
}
