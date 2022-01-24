package tests

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"store_server/internal/storage/mongorepo/dbconf"
	passportrepo "store_server/internal/storage/mongorepo/passport"
	routerepo "store_server/internal/storage/mongorepo/route"
)

const (
	dbName               = "viks_test"
	changeDateCollection = "changeDate_test"
	passportsCollection  = "passports_test"
	routesCollection     = "routes_test"
)

type client struct {
	passportRepository passportrepo.PassportRepository
	routeRepository    routerepo.RouteRepository
	db                 *mongo.Client
}

func (c *client) PassportRepository() passportrepo.PassportRepository {
	return c.passportRepository
}

func (c *client) RouteRepository() routerepo.RouteRepository {
	return c.routeRepository
}

func newClient() (*client, error) {
	path := "mongodb://localhost:27017/"
	clientOptions := options.Client().ApplyURI(path)
	cl, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = cl.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	conf := dbconf.DbConf{
		DatabaseName:             dbName,
		ChangeDateCollectionName: changeDateCollection,
		PassportsCollectionName:  passportsCollection,
		RoutesCollectionName:     routesCollection,
	}

	rc := client{
		passportRepository: passportrepo.NewPassportRepository(cl, conf),
		routeRepository:    routerepo.NewRouteRepository(cl, conf),
		db:                 cl,
	}
	return &rc, err
}
