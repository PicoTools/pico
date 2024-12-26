package utils

import (
	"os"
)

// EnvOr if environment variable is not specified returns default value
func EnvOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
