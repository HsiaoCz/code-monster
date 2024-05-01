package config

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return ":9001"
	}
	return port
}
