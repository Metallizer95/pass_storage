package routerepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
	"store_server/internal/storage/mongorepo/dbconf"
	passportrepo "store_server/internal/storage/mongorepo/passport"
	"store_server/pkg/logging"
)

type RouteRepository interface {
	Create(route routers.ViksRoute) *routers.ViksRoute
	Read(id string) *routers.ViksRoute
	ReadAll() []routers.ViksRoute
	Update(passport routers.ViksRoute) *routers.ViksRoute
	Delete(route routers.ViksRoute) *routers.ViksRoute
}

type routeRepositoryImpl struct {
	client             *mongo.Client
	logger             *logging.Logger
	routeCollection    *mongo.Collection
	passportCollection *mongo.Collection
}

func NewRouteRepository(db *mongo.Client, conf dbconf.DbConf) RouteRepository {
	logger, err := logging.GetLogger()
	if err != nil {
		panic(err)
	}

	routeCollection := db.Database(conf.DatabaseName).Collection(conf.RoutesCollectionName)
	passportCollection := db.Database(conf.DatabaseName).Collection(conf.PassportsCollectionName)

	return &routeRepositoryImpl{
		client:             db,
		logger:             logger,
		routeCollection:    routeCollection,
		passportCollection: passportCollection,
	}
}

func (r *routeRepositoryImpl) Create(route routers.ViksRoute) *routers.ViksRoute {
	model := routeToRepositoryModel(route)
	_, err := r.routeCollection.InsertOne(context.TODO(), model)
	if err != nil {
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

	passportsModel, err := r.findRoutePassports(route.Route.SectionSet)
	if err == nil && len(passportsModel) > 0 {
		var passports []passport.Passport
		for _, p := range passportsModel {
			passports = append(passports, passportrepo.ModelToPassport(p))
		}
		result.SectionSet = passports
	}

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
