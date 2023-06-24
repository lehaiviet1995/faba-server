package config

import "os"

var (
	ServerPort string
)

func init() {
	ServerPort = GetEnv("SERVER_PORT", ":8080")
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
