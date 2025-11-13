package configlib

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type BaseConfig struct {
	Service  ServiceConfig  `mapstructure:"service"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServiceConfig struct {
	Name        string `mapstructure:"name" validate:"required"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment" validate:"oneof=dev stage prod"`
}

type ServerConfig struct {
	GRPC GRPCConfig `mapstructure:"grpc"`
	HTTP HTTPConfig `mapstructure:"http"`
}

type GRPCConfig struct {
	Port int `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type HTTPConfig struct {
	Port int `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

// Load загружает конфигурацию из YAML файла
func Load() (*BaseConfig, error) {
	v := viper.New()

	// Настройки по умолчанию
	setDefaults(v)

	// Загружаем конфиг из файла
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file : %w", err)
	}

	// Биндим на структуру
	var cfg BaseConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Валидация
	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}

// MustLoad загружает конфиг или паникает
func MustLoad() *BaseConfig {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("service.environment", "dev")
	v.SetDefault("service.version", "1.0.0")

	v.SetDefault("server.grpc.port", 50051)
	v.SetDefault("server.http.port", 8080)

	v.SetDefault("database.port", 5432)
	v.SetDefault("database.ssl_mode", "disable")

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func validateConfig(cfg *BaseConfig) error {
	validate := validator.New()
	return validate.Struct(cfg)
}

// Хелперы для удобства использования конфига
func (c *BaseConfig) IsDev() bool {
	return c.Service.Environment == "dev"
}

func (c *BaseConfig) IsProd() bool {
	return c.Service.Environment == "prod"
}

func (c *BaseConfig) IsStage() bool {
	return c.Service.Environment == "stage"
}

func (c *BaseConfig) GetDBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name)

	// return fmt.Sprintf(
	// 	"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
	// 	c.Database.Host,
	// 	c.Database.Port,
	// 	c.Database.Name,
	// 	c.Database.Username,
	// 	c.Database.Password,
	// 	c.Database.SSLMode,
	// )
}

func (c *BaseConfig) GetGRPCAddress() string {
	return fmt.Sprintf(":%d", c.Server.GRPC.Port)
}

func (c *BaseConfig) GetHTTPAddress() string {
	return fmt.Sprintf(":%d", c.Server.HTTP.Port)
}

// Validate проверяет обязательные поля
func (c *BaseConfig) Validate() error {
	if c.Service.Name == "" {
		return fmt.Errorf("service name is required")
	}
	if c.Server.GRPC.Port == 0 {
		return fmt.Errorf("gRPC port is required")
	}
	if c.Database.Host == "" && c.Service.Name != "gateway" {
		return fmt.Errorf("database host is required for service %s", c.Service.Name)
	}
	return nil
}
