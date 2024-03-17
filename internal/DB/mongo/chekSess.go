package mongo

import (
	"context"
	"medods/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Mongo) ChekSess(reqId string, token string) (*model.Session, error) {
	var sess model.Session
	filter := bson.D{primitive.E{Key: "bcryptTocken", Value: token}}
	res := m.collection.FindOne(context.TODO(), filter)

	err := res.Decode(&sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}
