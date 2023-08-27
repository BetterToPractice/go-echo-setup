package lib

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
)

var configPath = "./config.yml"

var defaultConfig = Config{
	Name: "go-echo-setup",
	Http: &HttpConfig{
		Host: "0.0.0.0",
		Port: 9999,
	},
	Log: &LogConfig{},
	Database: &DatabaseConfig{
		Parameters:   "",
		MaxLifetime:  7200,
		MaxOpenConns: 150,
		MaxIdleConns: 50,
	},
	Cors: &CorsConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	},
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
	Name     string          `mapstructure:"Name"`
	Http     *HttpConfig     `mapstructure:"Http"`
	Log      *LogConfig      `mapstructure:"Log"`
	Database *DatabaseConfig `mapstructure:"Database"`
	Cors     *CorsConfig     `mapstructure:"Cors"`
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

type DatabaseConfig struct {
	Engine      string `mapstructure:"Engine"`
	Name        string `mapstructure:"Name"`
	Host        string `mapstructure:"Host"`
	Port        int    `mapstructure:"Port"`
	Username    string `mapstructure:"Username"`
	Password    string `mapstructure:"Password"`
	TablePrefix string `mapstructure:"TablePrefix"`
	Parameters  string `mapstructure:"Parameters"`
	SslMode     string `mapstructure:"SslMode"`
	TimeZone    string `mapstructure:"TimeZone"`

	MaxLifetime  int `mapstructure:"MaxLifetime"`
	MaxOpenConns int `mapstructure:"MaxOpenConns"`
	MaxIdleConns int `mapstructure:"MaxIdleConns"`
}

type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"AllowOrigins"`
	AllowMethods     []string `mapstructure:"AllowMethods"`
	AllowHeaders     []string `mapstructure:"AllowHeaders"`
	AllowCredentials bool     `mapstructure:"AllowCredentials"`
}

func (a DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", a.Host, a.Username, a.Password, a.Name, a.Port, a.SslMode, a.TimeZone)
}

func (a *HttpConfig) ListenAddr() string {
	if err := validator.New().Struct(a); err != nil {
		return "0.0.0.0:5111"
	}

	host := a.Host
	port := a.Port
	if host == "localhost" {
		host = "127.0.0.1"
	}

	return fmt.Sprintf("%s:%d", host, port)
}

func SetConfigPath(path string) {
	configPath = path
}
