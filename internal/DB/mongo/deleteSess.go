package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Удаляет сессию соответствующую строке token
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

// сверяет сессии в базе данных с текущим временем и удаляет если они устарели
func (m *Mongo) DeleteOld() {
	times := time.Now().Unix()

	filter := bson.D{primitive.E{Key: "expiretime", Value: bson.D{primitive.E{Key: "$lt", Value: times}}}}

	res, err := m.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		m.logger.L.WithField("Mongo.DeleteOne", "").Error(err)
	}

	if res.DeletedCount == 0 {
		m.logger.L.WithField("Mongo.DeleteOne", "Нет документов для удаления").Debug(times)
	} else {
		m.logger.L.WithField("Mongo.DeleteOne", fmt.Sprintf("Удалено документов %d", res.DeletedCount)).Debug()
	}

}
