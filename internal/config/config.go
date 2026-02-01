package config

import (
	"fmt"
	"log"
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
	// Загружаем .env — если файла нет, это НЕ ошибка
	if err := godotenv.Load(); err != nil {
		// Проверяем: это "файл не найден" или реальная ошибка?
		if !os.IsNotExist(err) {
			// Например, файл есть, но повреждён — это уже проблема
			log.Printf("Warning: failed to load .env: %v", err)
		} else {
			log.Println("Info: .env file not found, using system environment")
		}
	}

	// 2. Настраиваем Viper
	viper.SetConfigType("env") // говорим, что источник — env-стиль
	viper.AutomaticEnv()       // разрешаем брать из системного окружения

	// 3. Регистрируем все ключи, которые хотим прочитать
	keys := []string{
		"APP_NAME", "HTTP_HOST", "HTTP_PORT", "GRPC_PORT", "LOG_LEVEL", "DEBUG_MODE",
		"DB_HOST", "DB_PORT", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_USE_TLS",
		"REDIS_HOST", "REDIS_PREFIX", "KAFKA_HOST", "KAFKA_GROUP_ID",
	}

	// Viper не читает автоматически все env vars — нужно "подсказать"
	for _, key := range keys {
		viper.BindEnv(key)
	}

	// 4. Читаем конфиг в структуру
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	// 5. Строим PostgresDSN (как у тебя было)
	sslmode := "disable"
	if cfg.DBTLS {
		sslmode = "require"
	}

	// Экранируем user и password на случай спецсимволов
	user := url.QueryEscape(cfg.DBUser)
	password := url.QueryEscape(cfg.DBPassword)

	cfg.PostgresDSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		user, password, cfg.DBHost, cfg.DBPort, cfg.DBName, sslmode,
	)

	return &cfg
}
