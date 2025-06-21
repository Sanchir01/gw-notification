package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env    string    `yaml:"env"`
	Domain string    `yaml:"domain"`
	DB     Databases `yaml:"db"`
	Kafka  Kafka     `yaml:"kafka"`
}
type Databases struct {
	Mongo Mongo `yaml:"mongo"`
}

type Kafka struct {
	Consumer ConsumerConfig `yaml:"consumer"`
}

type ConsumerConfig struct {
	GroupID string   `yaml:"group_id"`
	Topic   []string `yaml:"topic"`
	Brokers []string `yaml:"brokers"`
}
type Mongo struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Dbname string `yaml:"database"`
}

func InitConfig() *Config {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env.dev"
	}
	fmt.Println("env name", envFile)
	if err := godotenv.Load(envFile); err != nil {
		slog.Error("ошибка при инициализации переменных окружения", err.Error())
	}
	configPath := os.Getenv("CONFIG_PATH")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist:%s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	return &cfg
}
