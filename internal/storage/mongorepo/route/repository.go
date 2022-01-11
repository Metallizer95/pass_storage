package routerepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/routers"
	"store_server/internal/storage/mongorepo/dbconf"
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

	routeCollection := db.Database(dbconf.DatabaseName).Collection(dbconf.RoutesCollectionName)

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
	route, ok := r.findRoute(id)
	if !ok {
		return nil
	}
	result := repositoryModelToRoute(*route)
	return &result
}

func (r *routeRepositoryImpl) ReadAll() []routers.ViksRoute {
	var resultRoutes []routers.ViksRoute
	routes := r.findAllRoutes()

	for _, route := range routes {
		resultRoutes = append(resultRoutes, repositoryModelToRoute(route))
	}

	return resultRoutes
}

func (r *routeRepositoryImpl) Update(route routers.ViksRoute) *routers.ViksRoute {
	model := routeToRepositoryModel(route)

	updatedModel, err := r.updateRoute(model)
	if err != nil {
		return nil
	}

	result := repositoryModelToRoute(*updatedModel)
	return &result
}

func (r *routeRepositoryImpl) Delete(route routers.ViksRoute) *routers.ViksRoute {
	model := routeToRepositoryModel(route)

	deletedModel, err := r.deleteRoute(model.ID)
	if err != nil {
		r.logger.Errorf("delete route error with id %s: %v", model.ID, err)
		return nil
	}

	result := repositoryModelToRoute(*deletedModel)
	return &result
}
