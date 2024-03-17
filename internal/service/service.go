package service

import (
	"medods/internal/DB/mongo"
	"medods/internal/logger"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	logger *logger.Logger
	db     *mongo.Mongo
}

func New(log *logger.Logger, db *mongo.Mongo) *Service {
	return &Service{
		logger: log,
		db:     db,
	}
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ClaimsRT struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
