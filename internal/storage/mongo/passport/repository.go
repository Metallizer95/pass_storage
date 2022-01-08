package passportrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/passport"
)

type PassportRepository interface {
	Create(passport passport.Data) *passport.Passport
	Read(id string) *passport.Passport
	ReadAll() []passport.Passport
	Update(passport passport.Passport) *passport.Passport
}

type passportRepositoryImpl struct {
	client mongo.Client
	cache  Cache
}

func NewPassportRepository(db mongo.Client) PassportRepository {
	return nil
}

func (m *passportRepositoryImpl) Create(passport passport.Data) *passport.Passport {
	return nil
}

func (m *passportRepositoryImpl) Read(id string) *passport.Passport {
	return nil
}

func (m *passportRepositoryImpl) ReadAll() []passport.Passport {
	return nil
}

func (m *passportRepositoryImpl) Update(passport passport.Passport) *passport.Passport {
	return nil
}

func (m *passportRepositoryImpl) Delete(id string) *passport.Passport {
	return nil
}
