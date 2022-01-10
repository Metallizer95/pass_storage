package passportrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pr *passportRepositoryImpl) findByIdPassport(id string) (*RepositoryModel, bool) {
	filter := bson.M{"_id": id}
	cursor, err := pr.passportCollections.Find(context.TODO(), filter)
	if err != nil {
		pr.logger.Errorf("search passport error: %v", err)
		return nil, false
	}

	var result []RepositoryModel
	if err := cursor.All(context.TODO(), &result); err != nil {
		pr.logger.Errorf("search passport error: %v", err)
		return nil, false
	}

	if len(result) > 0 {
		return &result[0], true
	}
	return nil, false
}

func (pr *passportRepositoryImpl) findAllPassports() []RepositoryModel {
	filter := bson.M{}
	cursor, err := pr.passportCollections.Find(context.TODO(), filter)
	if err != nil {
		pr.logger.Errorf("search passport error: %v", err)
		return nil
	}

	var result []RepositoryModel
	if err := cursor.All(context.TODO(), &result); err != nil {
		pr.logger.Errorf("search passport error with cursor: %v", err)
		return nil
	}

	return result
}

func (pr *passportRepositoryImpl) updatePassport(p RepositoryModel) error {
	filter := bson.D{primitive.E{Key: "_id", Value: p.ID}}
	result := pr.passportCollections.FindOneAndUpdate(context.TODO(), filter, p)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (pr *passportRepositoryImpl) deletePassport(id string) (*RepositoryModel, error) {
	filter := bson.M{"_id": id}
	singleResult := pr.passportCollections.FindOneAndDelete(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}
	var result RepositoryModel
	if err := singleResult.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
