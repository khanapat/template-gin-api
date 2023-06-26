package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string        `mapstructure:"name"`
		Port    string        `mapstructure:"port"`
		Timeout time.Duration `mapstructure:"timeout"`
		Context string        `mapstructure:"context"`
		Cors    struct {
			Origin string `mapstructure:"origin"`
		} `mapstructure:"cors"`
	} `mapstructure:"app"`
	Log struct {
		Level string `mapstructure:"level"`
		Env   string `mapstructure:"env"`
	} `mapstructure:"log"`
	Postgres struct {
		Type         string `mapstructure:"type"`
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		Username     string `mapstructure:"username"`
		Password     string `mapstructure:"password"`
		DBName       string `mapstructure:"database"`
		Timeout      int    `mapstructure:"timeout"`
		SSLMode      string `mapstructure:"sslmode"`
		PoolMaxConns int    `mapstructure:"pool-max-conns"`
	} `mapstructure:"postgres"`
	Redis struct {
		MaxIdle  int           `mapstructure:"max-idle"`
		Timeout  time.Duration `mapstructure:"timeout"`
		Host     string        `mapstructure:"host"`
		Password string        `mapstructure:"password"`
	} `mapstructure:"redis"`
	Client struct {
		Timeout  time.Duration `mapstructure:"timeout"`
		Hidebody bool          `mapstructure:"hidebody"`
	} `mapstructure:"client"`
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	viper.SetDefault("app.name", "template-gin-api")
	viper.SetDefault("app.port", "9090")
	viper.SetDefault("app.timeout", "60s")
	viper.SetDefault("app.context", "/api")
	viper.SetDefault("app.cors.origin", "*")

	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.env", "dev")

	viper.SetDefault("postgres.type", "postgres")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", "5432")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "")
	viper.SetDefault("postgres.database", "postgres")
	viper.SetDefault("postgres.timeout", 100)
	viper.SetDefault("postgres.sslmode", "disable")
	viper.SetDefault("postgres.pool-max-conns", 10)

	viper.SetDefault("redis.max-idle", 3)
	viper.SetDefault("redis.timeout", "60s")
	viper.SetDefault("redis.host", "localhost:6379")
	viper.SetDefault("redis.password", "")

	viper.SetDefault("client.timeout", "60s")
	viper.SetDefault("client.hidebody", true)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fmt.Println("--- configurations ---")
	for k, v := range viper.AllSettings() {
		fmt.Println(k, ":", v)
	}
	fmt.Println("----------------------")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
