package config

import (
	"os"
)

type configDb struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func DbNew() *configDb {

	return &configDb{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
