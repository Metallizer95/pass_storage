package routestorage

import (
	"store_server/internal/domain/routers"
)

type RepositoryInMemoryImpl struct {
	data map[string]routers.ViksRoute
}

func New() *RepositoryInMemoryImpl {
	return &RepositoryInMemoryImpl{data: make(map[string]routers.ViksRoute)}
}

func (r *RepositoryInMemoryImpl) Create(p routers.ViksRoute) *routers.ViksRoute {
	_, ok := r.data[p.ViksRoutedID]
	if ok {
		return nil
	}
	r.data[p.ViksRoutedID] = p
	return &p
}

func (r *RepositoryInMemoryImpl) Read(id string) *routers.ViksRoute {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	return &p
}

func (r *RepositoryInMemoryImpl) Update(route routers.ViksRoute) *routers.ViksRoute {
	_, ok := r.data[route.ViksRoutedID]
	if !ok {
		return nil
	}
	r.data[route.ViksRoutedID] = route
	result := r.data[route.ViksRoutedID]
	return &result
}

func (r *RepositoryInMemoryImpl) Delete(route routers.ViksRoute) *routers.ViksRoute {
	p, ok := r.data[route.ViksRoutedID]
	if !ok {
		return nil
	}
	delete(r.data, route.ViksRoutedID)
	return &p
}

func (r *RepositoryInMemoryImpl) ReadAll() []routers.ViksRoute {
	var routes []routers.ViksRoute
	for _, v := range r.data {
		routes = append(routes, v)
	}
	return routes
}
