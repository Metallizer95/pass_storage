package passportrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (pr *passportRepositoryImpl) FindByIdPassportCollection(id string) bool {
	filter := bson.M{"id": id}
	cursor, err := pr.passportCollections.Find(context.TODO(), filter)
	if err != nil {
		pr.logger.Errorf("search passport error: %v", err)
		return false
	}

	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		pr.logger.Errorf("search passport error: %v", err)
		return false
	}

	if len(result) > 0 {
		return true
	}
	return false
}

func (pr *passportRepositoryImpl) FindByIdChangeDateCollection(id string) bool {
	filter := bson.M{"_id": id}
	cursor, err := pr.changeDateCollection.Find(context.TODO(), filter)
	if err != nil {
		pr.logger.Errorf("search change date error: %v", err)
		return false
	}

	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		pr.logger.Errorf("search change date error: %v", err)
		return false
	}

	if len(result) > 0 {
		return true
	}
	return false
}
