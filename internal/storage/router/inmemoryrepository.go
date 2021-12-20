package storagerouter

import (
	"store_server/internal/domain/routers"
)

type RepositoryInMemoryImpl struct {
	data      map[string]routers.Router
	idCounter int
}

func New() *RepositoryInMemoryImpl {
	return &RepositoryInMemoryImpl{data: make(map[string]routers.Router), idCounter: 0}
}

func (r *RepositoryInMemoryImpl) Create(p routers.Router) *routers.Router {
	r.data[p.ID] = p
	return &p
}
func (r *RepositoryInMemoryImpl) Read(id string) *routers.Router {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	return &p
}

func (r *RepositoryInMemoryImpl) Update(route routers.Router) *routers.Router {
	_, ok := r.data[route.ID]
	if !ok {
		return nil
	}
	r.data[route.ID] = route
	result := r.data[route.ID]
	return &result
}

func (r *RepositoryInMemoryImpl) Delete(id string) *routers.Router {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	delete(r.data, id)
	return &p
}
