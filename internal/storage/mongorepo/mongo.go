package mongorepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	passportrepo "store_server/internal/storage/mongorepo/passport"
	routerepo "store_server/internal/storage/mongorepo/route"
)

type Config struct {
	Path string `json:"path"`
}

type Client interface {
	PassportRepository() passportrepo.PassportRepository
	RouteRepository() routerepo.RouteRepository
}

type repoClient struct {
	passportRepository passportrepo.PassportRepository
	routeRepository    routerepo.RouteRepository
}

func NewClient(cfg *Config) (Client, error) {
	if cfg == nil {
		cfg = &Config{Path: "mongodb://localhost:27017/"}
	}
	clientOptions := options.Client().ApplyURI(cfg.Path)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	rc := repoClient{
		passportRepository: passportrepo.NewPassportRepository(client),
		routeRepository:    routerepo.NewRouteRepository(client),
	}
	return &rc, err
}

func (c *repoClient) PassportRepository() passportrepo.PassportRepository {
	return c.passportRepository
}

func (c *repoClient) RouteRepository() routerepo.RouteRepository {
	return c.routeRepository
}
