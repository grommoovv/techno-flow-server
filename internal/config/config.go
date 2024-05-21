package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Server   ServerConfig
		Postgres PostgresConfig
		Jwt      JwtConfig
	}

	ServerConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	PostgresConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
		SslMode  string `mapstructure:"sslmode"`
	}

	JwtConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	}
)

func Init() (*Config, error) {

	if err := readConfigFile(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	var config Config

	if err := unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func readConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("development")

	return viper.ReadInConfig()
}

func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("server", &config.Server); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &config.Postgres); err != nil {
		return err
	}

	return viper.UnmarshalKey("jwt", &config.Jwt)
}
