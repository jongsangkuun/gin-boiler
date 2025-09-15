package config

import (
	"os"
)

var ENV *Env

type Env struct {
	ApiConfig
	DbConfig
	DbConnectionPool
}

type ApiConfig struct {
	Port         string
	AllowOrigins []string
	AllowMethods []string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}
type DbConnectionPool struct {
	MaxIdleConns    string
	MaxOpenConns    string
	ConnMaxLifetime string
	ConnMaxIdleTime string
}

func LoadEnv() (Env, error) {
	envStatus := os.Getenv("ENV_STATUS")
	if envStatus == "local" {
		dbConf := DbConfig{
			Host:     os.Getenv("DEV_POSTGRES_HOST"),
			Port:     os.Getenv("DEV_POSTGRES_PORT"),
			User:     os.Getenv("DEV_POSTGRES_USER"),
			Password: os.Getenv("DEV_POSTGRES_PASSWORD"),
			DbName:   os.Getenv("DEV_POSTGRES_DB"),
			SslMode:  os.Getenv("DEV_POSTGRES_SSL_MODE"),
		}
		dbConnPool := DbConnectionPool{
			MaxIdleConns:    os.Getenv("DEV_POSTGRES_MAX_IDLE_CONNS"),
			MaxOpenConns:    os.Getenv("DEV_POSTGRES_MAX_OPEN_CONNS"),
			ConnMaxLifetime: os.Getenv("DEV_POSTGRES_CONN_MAX_LIFETIME"),
			ConnMaxIdleTime: os.Getenv("DEV_POSTGRES_CONN_MAX_IDLE_TIME"),
		}

		return Env{
			ApiConfig{},
			dbConf,
			dbConnPool,
		}, nil
	} else {
		dbConf := DbConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DbName:   os.Getenv("POSTGRES_DB"),
			SslMode:  os.Getenv("POSTGRES_SSL_MODE"),
		}
		dbConnPool := DbConnectionPool{
			MaxIdleConns:    os.Getenv("POSTGRES_MAX_IDLE_CONNS"),
			MaxOpenConns:    os.Getenv("POSTGRES_MAX_OPEN_CONNS"),
			ConnMaxLifetime: os.Getenv("POSTGRES_CONN_MAX_LIFETIME"),
			ConnMaxIdleTime: os.Getenv("POSTGRES_CONN_MAX_IDLE_TIME"),
		}

		return Env{
			ApiConfig{},
			dbConf,
			dbConnPool,
		}, nil
	}
}
