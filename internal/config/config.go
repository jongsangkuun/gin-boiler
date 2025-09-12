package config

import (
	"os"
)

type Env struct {
	apiConfig
	dbConfig
	dbConnectionPool
}

type apiConfig struct {
	port         string
	allowOrigins []string
	allowMethods []string
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}
type dbConnectionPool struct {
	maxIdleConns    string
	maxOpenConns    string
	connMaxLifetime string
	connMaxIdleTime string
}

func LoadEnv() (Env, error) {
	envStatus := os.Getenv("ENV_STATUS")
	if envStatus == "local" {
		dbConf := dbConfig{
			host:     os.Getenv("DEV_POSTGRES_HOST"),
			port:     os.Getenv("DEV_POSTGRES_PORT"),
			user:     os.Getenv("DEV_POSTGRES_USER"),
			password: os.Getenv("DEV_POSTGRES_PASSWORD"),
			dbName:   os.Getenv("DEV_POSTGRES_NAME"),
			sslMode:  os.Getenv("DEV_POSTGRES_SSL_MODE"),
		}
		dbConnPool := dbConnectionPool{
			maxIdleConns:    os.Getenv("DEV_POSTGRES_MAX_IDLE_CONNS"),
			maxOpenConns:    os.Getenv("DEV_POSTGRES_MAX_OPEN_CONNS"),
			connMaxLifetime: os.Getenv("DEV_POSTGRES_CONN_MAX_LIFETIME"),
			connMaxIdleTime: os.Getenv("DEV_POSTGRES_CONN_MAX_IDLE_TIME"),
		}

		return Env{
			apiConfig{},
			dbConf,
			dbConnPool,
		}, nil
	} else {
		dbConf := dbConfig{
			host:     os.Getenv("POSTGRES_HOST"),
			port:     os.Getenv("POSTGRES_PORT"),
			user:     os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASSWORD"),
			dbName:   os.Getenv("POSTGRES_NAME"),
			sslMode:  os.Getenv("POSTGRES_SSL_MODE"),
		}
		dbConnPool := dbConnectionPool{
			maxIdleConns:    os.Getenv("POSTGRES_MAX_IDLE_CONNS"),
			maxOpenConns:    os.Getenv("POSTGRES_MAX_OPEN_CONNS"),
			connMaxLifetime: os.Getenv("POSTGRES_CONN_MAX_LIFETIME"),
			connMaxIdleTime: os.Getenv("POSTGRES_CONN_MAX_IDLE_TIME"),
		}

		return Env{
			apiConfig{},
			dbConf,
			dbConnPool,
		}, nil
	}
}
