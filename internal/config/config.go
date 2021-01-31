package config

import "os"

// Config ...
type Config struct {
	Environment string
	ConnectionStr string
}

func getConfigValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}

	return defaultValue
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Environment: getConfigValue("ENV", "local"),
		ConnectionStr: getConfigValue("CONN_STRING", "root:root@tcp(localhost:3306)/votingdb"),
	}
}