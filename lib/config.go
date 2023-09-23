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
	Name:   "go-echo-setup",
	Secret: "foobar",
	Http: &HttpConfig{
		Host: "0.0.0.0",
		Port: 9999,
	},
	Log: &LogConfig{},
	Database: &DatabaseConfig{
		Parameters:   "",
		MigrationDir: "migrations",
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
	Swagger: &SwaggerConfig{
		Title:       "Go Echo Setup Docs",
		Description: "Collection of Endpoints",
		Version:     "1.0",
		DocUrl:      "/swagger/*",
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
	Secret   string          `mapstructure:"Secret"`
	Http     *HttpConfig     `mapstructure:"Http"`
	Log      *LogConfig      `mapstructure:"Log"`
	Database *DatabaseConfig `mapstructure:"Database"`
	Cors     *CorsConfig     `mapstructure:"Cors"`
	Swagger  *SwaggerConfig  `mapstructure:"Swagger"`
	Auth     *AuthConfig     `mapstructure:"Auth"`
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
	Engine       string `mapstructure:"Engine"`
	Name         string `mapstructure:"Name"`
	Host         string `mapstructure:"Host"`
	Port         int    `mapstructure:"Port"`
	Username     string `mapstructure:"Username"`
	Password     string `mapstructure:"Password"`
	Parameters   string `mapstructure:"Parameters"`
	SslMode      string `mapstructure:"SslMode"`
	TimeZone     string `mapstructure:"TimeZone"`
	MigrationDir string `mapstructure:"MigrationDir"`
}

type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"AllowOrigins"`
	AllowMethods     []string `mapstructure:"AllowMethods"`
	AllowHeaders     []string `mapstructure:"AllowHeaders"`
	AllowCredentials bool     `mapstructure:"AllowCredentials"`
}

type SwaggerConfig struct {
	Title       string `mapstructrue:"Title"`
	Description string `mapstructure:"Description"`
	Version     string `mapstructure:"Version"`
	DocUrl      string `mapstructure:"DocUrl"`
}

type AuthConfig struct {
	Enable             string   `mapstructure:"Enable"`
	TokenExpired       int      `mapstructure:"TokenExpired"`
	IgnorePathPrefixes []string `mapstructure:"IgnorePathPrefixes"`
}

func (a DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", a.Host, a.Username, a.Password, a.Name, a.Port, a.SslMode, a.TimeZone)
}

func (a *HttpConfig) ListenAddr() string {
	if err := validator.New().Struct(a); err != nil {
		return "0.0.0.0:5000"
	}
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func SetConfigPath(path string) {
	configPath = path
}
