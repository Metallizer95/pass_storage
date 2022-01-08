package passportrepo

import (
	"store_server/internal/domain/passport"
)

type MongoPassportRepository struct {
}

func (m *MongoPassportRepository) Create(passport passport.Data) *passport.Passport {
	return nil
}

func (m *MongoPassportRepository) Read(id string) *passport.Passport {
	return nil
}

func (m *MongoPassportRepository) ReadAll() []passport.Passport {
	return nil
}

func (m *MongoPassportRepository) Update(passport passport.Passport) *passport.Passport {
	return nil
}

func (m *MongoPassportRepository) Delete(id string) *passport.Passport {
	return nil
}
