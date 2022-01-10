package routerepo

import "store_server/internal/domain/routers"

type RouteRepository interface {
	Create(route routers.ViksRoute) *routers.ViksRoute
	Read(id string) *routers.ViksRoute
	ReadAll() []routers.ViksRoute
	Update(passport routers.ViksRoute) *routers.ViksRoute
	Delete(route routers.ViksRoute) *routers.ViksRoute
}

type routeRepositoryImpl struct {
}

func (r *routeRepositoryImpl) Create(route routers.ViksRoute) *routers.ViksRoute {
	return nil
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
