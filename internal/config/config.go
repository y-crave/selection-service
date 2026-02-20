package config

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"

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

func Load() (*Config, error) {
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
	setDefaults()
	// 2. Настраиваем Viper — достаточно одного вызова!
	viper.AutomaticEnv() // автоматически ищет переменные в UPPER_SNAKE_CASE (как в наших тегах)

	// 3. Читаем напрямую в структуру — без ручного списка ключей!
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error("failed to unmarshal config", "config", err)
		return nil, errors.ErrUnmarshal
	}
	// 4. Валидация ОБЯЗАТЕЛЬНЫХ полей (Эти поля НЕ имеют дефолтов и ДОЛЖНЫ быть заданы в окружении)
	if err := validateRequired(&cfg); err != nil {
		log.Error("configuration validation failed", "error", err)
		return nil, errors.ErrValidRequired
	}
	// 5. Строим DSN (POSTGRES_DSN не читаем из окружения — генерируем сами)
	sslmode := "disable"
	if cfg.DBTLS {
		sslmode = "require"
	}
	user := url.QueryEscape(cfg.DBUser)
	password := url.QueryEscape(cfg.DBPassword)
	cfg.PostgresDSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s", user, password, cfg.DBHost, cfg.DBPort, cfg.DBName, sslmode,
	)

	return &cfg, nil
}

func setDefaults() {
	// === Приложение ===
	viper.SetDefault("APP_NAME", "base-service")
	viper.SetDefault("HTTP_HOST", "0.0.0.0") // Слушать все интерфейсы
	viper.SetDefault("HTTP_PORT", 8080)      // Стандартный HTTP порт
	viper.SetDefault("GRPC_PORT", 9090)      // Стандартный gRPC порт
	viper.SetDefault("LOG_LEVEL", "info")    // Уровень логирования
	viper.SetDefault("DEBUG_MODE", false)    // Режим отладки

	// === База данных (порт и TLS опциональны) ===
	viper.SetDefault("DB_PORT", 5432)     // Стандартный порт PostgreSQL
	viper.SetDefault("DB_USE_TLS", false) // Без TLS по умолчанию

	// === Redis (префикс опционален) ===
	viper.SetDefault("REDIS_PREFIX", "base-service")

	// === Kafka (group ID опционален) ===
	viper.SetDefault("KAFKA_GROUP_ID", "base-service-group")
}

func validateRequired(cfg *Config) error {
	var err []string
	if cfg.DBHost == "" {
		err = append(err, "DB_HOST is required (example: localhost, postgres, 10.0.0.1)")
	}
	if cfg.DBPassword == "" {
		err = append(err, "DB_PASSWORD is required (example: mysecretpassword)")
	}
	if cfg.RedisHost == "" {
		err = append(err, "REDIS_HOST is required (example: localhost:6379, redis:6379)")
	} else if !strings.Contains(cfg.RedisHost, ":") {
		err = append(err, fmt.Sprintf("❌ REDIS_HOST must include port (example: localhost:6379), got: %q", cfg.RedisHost))
	}
	if cfg.KafkaHost == "" {
		err = append(err, "KAFKA_HOST is required (example: localhost:9092, kafka:9092)")
	} else if !strings.Contains(cfg.KafkaHost, ":") {
		err = append(err, fmt.Sprintf("❌ KAFKA_HOST must include port (example: localhost:9092), got: %q", cfg.KafkaHost))
	}
	if len(err) > 0 {
		return fmt.Errorf(
			"configuration validation failed — these values MUST be set in environment variables or .env file:\n%s\n\n💡 Hint: Check your .env file or deployment configuration",
			strings.Join(err, "\n"),
		)
	}
	return nil
}
