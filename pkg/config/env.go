package config

import "os"

type MySQLInfo struct {
	MySQLUser     string
	MySQLPassword string
	MySQLPort     string
	MySQLHost     string
	MySQLDBName   string
}

type TokenInfo struct {
	SecretKey string
}

func LoadDBConfig() *MySQLInfo {
	mysqlUser := getEnv("MYSQL_USER", "docker")
	mysqlPassword := getEnv("MYSQL_PASSWORD", "docker")
	mysqlPort := getEnv("MYSQL_PORT", "3306")
	mysqlHost := getEnv("MYSQL_HOST", "db")
	mysqlDBName := getEnv("MYSQL_DATABASE", "test_database")
	return &MySQLInfo{
		MySQLUser:     mysqlUser,
		MySQLPassword: mysqlPassword,
		MySQLPort:     mysqlPort,
		MySQLHost:     mysqlHost,
		MySQLDBName:   mysqlDBName,
	}
}

func LoadTokenConfig() *TokenInfo {
	secretKey := getEnv("SECRET_KEY", "secret")
	return &TokenInfo{
		SecretKey: secretKey,
	}
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
