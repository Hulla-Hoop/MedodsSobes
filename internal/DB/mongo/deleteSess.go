package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

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
		return errors.New("нет документов для удаления")
	}

	return nil
}

func (m *Mongo) DeleteOld() {

	filter := bson.D{primitive.E{Key: "expiretime", Value: bson.D{primitive.E{Key: "$lt", Value: time.Now().Unix()}}}}

	res, err := m.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		m.logger.L.WithField("Mongo.DeleteOne", "").Error(err)
	}

	if res.DeletedCount == 0 {
		m.logger.L.WithField("Mongo.DeleteOne", "Нет документов для удаления").Debug()
	} else {
		m.logger.L.WithField("Mongo.DeleteOne", fmt.Sprintf("Удалено документов %d", res.DeletedCount)).Debug()
	}

}
