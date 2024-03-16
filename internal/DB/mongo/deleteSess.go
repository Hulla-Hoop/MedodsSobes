package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Mongo) DeleteSess(reqId string, token string) error {
	filter := bson.D{primitive.E{Key: "bcryptTocken", Value: token}}
	res, err := m.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no tasks were deleted")
	}

	return nil
}
