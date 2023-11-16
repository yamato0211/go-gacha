package config

import "os"

type MySQLInfo struct {
	MySQLUser     string
	MySQLPassword string
	MySQLPort     string
	MySQLHost     string
	MySQLDBName   string
}

type RedisInfo struct {
	Addr string
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

func LoadRedisConfig() *RedisInfo {
	redisAddr := getEnv("REDIS_ADDR", "redis:6379")
	return &RedisInfo{
		Addr: redisAddr,
	}
}

func getEnv(key string, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
