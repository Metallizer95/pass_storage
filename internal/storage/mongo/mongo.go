package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	passportrepo "store_server/internal/storage/mongo/passport"
	routerepo "store_server/internal/storage/mongo/route"
)

type Config struct {
	Path string `json:"path"`
}

type Client interface {
	PassportRepository() passportrepo.PassportRepository
	RouteRepository() routerepo.RepositoryModel
}

type client struct {
	passportRepository passportrepo.PassportRepository
	routeRepository    routerepo.RepositoryModel
}

func NewClient(cfg Config) error {
	clientOptions := options.Client().ApplyURI(cfg.Path)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) PassportRepository() passportrepo.PassportRepository {
	return c.passportRepository
}

func (c *client) RouteRepository() routerepo.RepositoryModel {
	return c.routeRepository
}
