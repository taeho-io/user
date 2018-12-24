package server

import "os"

type Settings struct {
	PostgresDBName   string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func NewSettings() Settings {
	return Settings{
		PostgresDBName:   getEnv("USER_POSTGRES_DB_NAME", "taeho"),
		PostgresHost:     getEnv("USER_POSTGRES_HOST", "127.0.0.1"),
		PostgresUser:     getEnv("USER_POSTGRES_USER", "taeho"),
		PostgresPassword: getEnv("USER_POSTGRES_PASSWORD", "WRONG_PASSWORD"),
	}
}

func MockSettings() Settings {
	return NewSettings()
}
