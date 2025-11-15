package configs

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type BaseConfig struct {
	Service  ServiceConfig  `mapstructure:"service"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServiceConfig struct {
	Title       string `mapstructure:"title"`
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

func Load() (*BaseConfig, error) {

	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("load .env file: %w", err)
	}

	v := viper.New()
	setDefaults(v)
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.BindEnv("database.host", "POSTGRES_HOST")
	v.BindEnv("database.port", "POSTGRES_PORT")
	v.BindEnv("database.name", "POSTGRES_DB")
	v.BindEnv("database.username", "POSTGRES_USER")
	v.BindEnv("database.password", "POSTGRES_PASSWORD")

	var cfg BaseConfig
	if err := v.Unmarshal(&cfg, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return &cfg, nil
}
func MustLoad() *BaseConfig {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("config load failed: %v", err))
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

func validateConfig(cfg *BaseConfig) error {
	validate := validator.New()
	return validate.Struct(cfg)
}

// === Хелперы ===
func (c *BaseConfig) IsDev() bool   { return c.Service.Environment == "dev" }
func (c *BaseConfig) IsProd() bool  { return c.Service.Environment == "prod" }
func (c *BaseConfig) IsStage() bool { return c.Service.Environment == "stage" }

func (c *BaseConfig) GetDBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *BaseConfig) GetGRPCAddress() string { return fmt.Sprintf("%d", c.Server.GRPC.Port) }
func (c *BaseConfig) GetHTTPAddress() string { return fmt.Sprintf("%d", c.Server.HTTP.Port) }
