package config

import (
	"fmt"

	"net/url"
	"os"
	"strconv"
)

type Config struct {
	PostgresDSN         string
	AppName             string
	AppHost             string
	AppHttpPort         int
	AppGrpcPort         int
	LogLevel            string
	DebugMode           bool
	DBHost              string
	DBPort              int
	DBName              string
	DBUser              string
	DBPassword          string
	DBTLS               bool
	RedisHost           string
	RedisPref           string
	KafkaHost           string
	KafkaGroupID        string
	UserFilterMinAge    int
	UserFilterMaxAge    int
	UserFilterMinHeight int
	UserFilterMaxHeight int
}

func Load() *Config {
	cfg := &Config{
		AppName:             getEnv("APP_NAME", "selection-service"),
		AppHost:             getEnv("HTTP_HOST", "0.0.0.0"),
		AppHttpPort:         getEnvAsInt("HTTP_PORT", 8080),
		AppGrpcPort:         getEnvAsInt("GRPC_PORT", 8081),
		LogLevel:            getEnv("LOG_LEVEL", "info"),
		DebugMode:           getEnvAsBool("DEBUG_MODE", false),
		DBHost:              getEnv("DB_HOST", "localhost"),
		DBPort:              getEnvAsInt("DB_PORT", 5432),
		DBName:              getEnv("DB_NAME", "selection_service"),
		DBUser:              getEnv("DB_USER", "selection"),
		DBPassword:          getEnv("DB_PASSWORD", "selection"),
		DBTLS:               getEnvAsBool("DB_USE_TLS", false),
		RedisHost:           getEnv("REDIS_HOST", "localhost:6379"),
		RedisPref:           getEnv("REDIS_PREFIX", "selection_"),
		KafkaHost:           getEnv("KAFKA_HOST", "localhost:9092"),
		KafkaGroupID:        getEnv("KAFKA_GROUP_ID", "selection.all"),
		UserFilterMaxAge:    getEnvAsInt("MAX_AGE", 100),
		UserFilterMinAge:    getEnvAsInt("MIN_AGE", 18),
		UserFilterMaxHeight: getEnvAsInt("MAX_HEIGHT", 300),
		UserFilterMinHeight: getEnvAsInt("MIN_HEIGHT", 50),
	}

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

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		fmt.Printf("%s=%s\n", key, value)
		return value
	}
	fmt.Printf("%s=%s\n", key, defaultValue)
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultValue))
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, strconv.FormatBool(defaultValue))
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
