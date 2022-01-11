package routerepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/routers"
	"store_server/internal/storage/mongorepo"
	"store_server/pkg/logging"
	"strings"
)

type RouteRepository interface {
	Create(route routers.ViksRoute) *routers.ViksRoute
	Read(id string) *routers.ViksRoute
	ReadAll() []routers.ViksRoute
	Update(passport routers.ViksRoute) *routers.ViksRoute
	Delete(route routers.ViksRoute) *routers.ViksRoute
}

type routeRepositoryImpl struct {
	client          *mongo.Client
	logger          *logging.Logger
	routeCollection *mongo.Collection
}

func NewRouteRepository(db *mongo.Client) RouteRepository {
	logger, err := logging.GetLogger()
	if err != nil {
		panic(err)
	}

	routeCollection := db.Database(mongorepo.DatabaseName).Collection(mongorepo.RoutesCollectionName)

	return &routeRepositoryImpl{
		client:          db,
		logger:          logger,
		routeCollection: routeCollection,
	}
}

func (r *routeRepositoryImpl) Create(route routers.ViksRoute) *routers.ViksRoute {
	model := routeToRepositoryModel(route)

	_, err := r.routeCollection.InsertOne(context.TODO(), model)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		return nil
	}
	return &route
}

func (r *routeRepositoryImpl) Read(id string) *routers.ViksRoute {
	return nil
}

func (r *routeRepositoryImpl) ReadAll() []routers.ViksRoute {
	return nil
}

func (r *routeRepositoryImpl) Update(passport routers.ViksRoute) *routers.ViksRoute {
	return nil
}

func (r *routeRepositoryImpl) Delete(route routers.ViksRoute) *routers.ViksRoute {
	return nil
}
