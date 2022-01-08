package passportstorage

import (
	"store_server/internal/domain/passport"
	passportrepo "store_server/internal/storage/mongo/passport"
	"store_server/pkg/cache"
	"time"
)

type RepositoryInMemoryImpl struct {
	data  map[string]passport.Data
	cache passportrepo.Cache
}

func New() *RepositoryInMemoryImpl {
	defaultExpiration := 10 * time.Minute
	cleanupInterval := 10 * time.Minute
	return &RepositoryInMemoryImpl{
		data:  make(map[string]passport.Data),
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

func (r *RepositoryInMemoryImpl) Create(p passport.Data) *passport.Passport {
	_, ok := r.data[p.SectionID]
	if ok {
		return nil
	}
	r.data[p.SectionID] = p
	result := passport.Passport{
		ID:   p.SectionID,
		Data: p,
	}
	return &result
}

func (r *RepositoryInMemoryImpl) Read(id string) *passport.Passport {
	p, ok := r.data[id]
	if !ok {
		return nil
	}
	cachedData, ok := r.cache.Get(id)
	if ok {
		cd := cachedData.(passport.Passport)
		return &cd
	}

	result := &passport.Passport{
		ID:   id,
		Data: p,
	}
	r.cache.Set(id, result, time.Minute*20)
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

	_ = r.cache.Delete(id)

	return &passport.Passport{
		ID:   id,
		Data: p,
	}
}

func (r *RepositoryInMemoryImpl) ReadAll() []passport.Passport {
	var result []passport.Passport
	for k, v := range r.data {
		result = append(result, passport.Passport{
			ID:   k,
			Data: v,
		})
	}
	return result
}
