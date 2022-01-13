package routerepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	passportrepo "store_server/internal/storage/mongorepo/passport"
)

func (r *routeRepositoryImpl) findRoute(id string) (*repositoryModel, bool) {
	filter := bson.M{"id": id}
	cursor, err := r.routeCollection.Find(context.TODO(), filter)
	if err != nil {
		r.logger.Errorf("find route in repository error: %v", err)
		return nil, false
	}

	var result []repositoryModel
	if err := cursor.All(context.TODO(), &result); err != nil {
		r.logger.Errorf("find route in repository error: %v", err)
		return nil, false
	}
	if len(result) > 0 {
		return &result[0], true
	}

	return nil, false
}

func (r *routeRepositoryImpl) findAllRoutes() []repositoryModel {
	var result []repositoryModel
	filter := bson.M{}
	cursor, err := r.routeCollection.Find(context.TODO(), filter)
	if err != nil {
		r.logger.Errorf("find all routes in repository error: %v", err)
		return result
	}

	if err := cursor.All(context.TODO(), &result); err != nil {
		r.logger.Errorf("find all routes in repository error: %v", err)
		return result
	}
	return result
}

func (r *routeRepositoryImpl) updateRoute(p repositoryModel) (*repositoryModel, error) {
	filter := bson.M{"id": p.ID}
	singleResult := r.routeCollection.FindOneAndUpdate(context.TODO(), filter, p)

	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	var result *repositoryModel
	if err := singleResult.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *routeRepositoryImpl) deleteRoute(id string) (*repositoryModel, error) {
	filter := bson.M{"id": id}

	singleResult := r.routeCollection.FindOneAndDelete(context.TODO(), filter)
	if singleResult.Err() != nil {
		return nil, singleResult.Err()
	}

	var result repositoryModel
	if err := singleResult.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *routeRepositoryImpl) findRoutePassports(ids []string) ([]passportrepo.RepositoryModel, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}

	cursor, err := r.passportCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var result []passportrepo.RepositoryModel
	if err := cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return result, nil
}
