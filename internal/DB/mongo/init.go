package mongo

import (
	"context"
	"medos/internal/config"
	"medos/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	collection *mongo.Collection
	logger     *logger.Logger
}

func New(log *logger.Logger) *Mongo {
	cfg := config.DbNew()
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.Host + ":" + cfg.Port + "/").SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.L.Info("не удалось подключиться к Mongo")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.L.Info("Mongo не доступна по протоколу IP")
	}
	collection := client.Database(cfg.DBName).Collection(cfg.DBName)

	log.L.Info("Mongo поднялось")

	return &Mongo{
		collection: collection,
		logger:     log,
	}
}
