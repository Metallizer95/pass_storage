package common

import (
	"store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
)

type Repository interface {
	RouteRepository() routers.Repository
	PassportRepository() passport.Repository
	FindPassportsByRoute(string) PassportsRoute
}
