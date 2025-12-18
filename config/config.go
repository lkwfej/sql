package config

import "os"

// DBConfig collects database connection parameters.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// LoadDBConfig loads database settings from environment variables with sane defaults
// so the application can run without manual edits when paired with the included
// docker-compose MySQL service.
func LoadDBConfig() DBConfig {
	return DBConfig{
		User:     getEnv("DB_USER", "admin_sql"),
		Password: getEnv("DB_PASSWORD", "admin_sql"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "3306"),
		Name:     getEnv("DB_NAME", "goland_demo"),
	}
}

// DSN renders the connection string for GORM/MySQL.
func (c DBConfig) DSN() string {
	return c.User + ":" + c.Password + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
