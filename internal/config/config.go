package config

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	PostgresDSN  string `mapstructure:"POSTGRES_DSN"` // можно не задавать — мы строим сами
	AppName      string `mapstructure:"APP_NAME"`
	AppHost      string `mapstructure:"HTTP_HOST"`
	AppHttpPort  int    `mapstructure:"HTTP_PORT"`
	AppGrpcPort  int    `mapstructure:"GRPC_PORT"`
	LogLevel     string `mapstructure:"LOG_LEVEL"`
	DebugMode    bool   `mapstructure:"DEBUG_MODE"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       int    `mapstructure:"DB_PORT"`
	DBName       string `mapstructure:"DB_NAME"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBTLS        bool   `mapstructure:"DB_USE_TLS"`
	RedisHost    string `mapstructure:"REDIS_HOST"`
	RedisPref    string `mapstructure:"REDIS_PREFIX"`
	KafkaHost    string `mapstructure:"KAFKA_HOST"`
	KafkaGroupID string `mapstructure:"KAFKA_GROUP_ID"`
}

func Load() *Config {
	log := slog.Default()
	// Загружаем .env — если файла нет, это НЕ ошибка
	if err := godotenv.Load(); err != nil {
		// Проверяем: это "файл не найден" или реальная ошибка?
		if !os.IsNotExist(err) {
			// Например, файл есть, но повреждён — это уже проблема
			log.Error("failed to load", ".env:", err)
		} else {
			log.Warn(".env file not found using system environment")
		}
	}

	// 2. Настраиваем Viper — достаточно одного вызова!
	viper.AutomaticEnv() // автоматически ищет переменные в UPPER_SNAKE_CASE (как в наших тегах)

	// 3. Читаем напрямую в структуру — без ручного списка ключей!
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error("failed to unmarshal config", "config", err)
		return nil
	}

	// 4. Строим DSN (POSTGRES_DSN не читаем из окружения — генерируем сами)
	sslmode := "disable"
	if cfg.DBTLS {
		sslmode = "require"
	}
	user := url.QueryEscape(cfg.DBUser)
	password := url.QueryEscape(cfg.DBPassword)
	cfg.PostgresDSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, cfg.DBHost, cfg.DBPort, cfg.DBName, sslmode,
	)

	return &cfg
}
