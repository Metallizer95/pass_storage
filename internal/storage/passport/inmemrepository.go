package passportstorage

import (
	"store_server/internal/domain/passport"
	"strconv"
)

type RepositoryInMemoryImpl struct {
	data      map[string]passport.Data
	idCounter int
}

func New() *RepositoryInMemoryImpl {
	return &RepositoryInMemoryImpl{data: make(map[string]passport.Data), idCounter: 0}
}

func (r *RepositoryInMemoryImpl) Create(p passport.Data) *passport.Passport {
	idstr := strconv.Itoa(r.idCounter)
	defer func() { r.idCounter++ }()
	r.data[idstr] = p
	result := passport.Passport{
		ID:   idstr,
		Data: p,
	}
	return &result
}
func (r *RepositoryInMemoryImpl) Read(id string) *passport.Passport {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	result := &passport.Passport{
		ID:   id,
		Data: p,
	}
	return result
}

func (r *RepositoryInMemoryImpl) Update(passport passport.Passport) *passport.Passport {
	_, ok := r.data[passport.ID]
	if !ok {
		return nil
	}
	r.data[passport.ID] = passport.Data
	return &passport
}

func (r *RepositoryInMemoryImpl) Delete(id string) *passport.Passport {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	delete(r.data, id)
	return &passport.Passport{
		ID:   id,
		Data: p,
	}
}
