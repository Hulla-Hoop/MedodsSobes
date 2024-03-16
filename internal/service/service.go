package service

import (
	"medos/internal/DB/mongo"
	"medos/internal/logger"

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

var users = map[string]string{
	"3825c945-8843-4b7d-995e-30b16c173c65": "user1",
	"019ed7ca-8286-40b8-ac80-1950c92dccfd": "user2",
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ClaimsRT struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
