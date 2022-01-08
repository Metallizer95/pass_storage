package passportrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/passport"
	"store_server/pkg/cache"
	"time"
)

const (
	DatabaseName             = "passport"
	changeDateCollectionName = "changeDate"
	passportsCollectionName  = "passports"
)

type PassportRepository interface {
	Create(passport passport.Data) *passport.Passport
	Read(id string) *passport.Passport
	ReadAll() []passport.Passport
	Update(passport passport.Passport) *passport.Passport
}

type passportRepositoryImpl struct {
	client               mongo.Client
	cache                Cache
	changeDateCollection *mongo.Collection
	passportCollections  *mongo.Collection
}

func NewPassportRepository(db mongo.Client) PassportRepository {
	cacheExpirationTime := 10 * time.Minute
	cacheCleanUpTime := 10 * time.Minute
	changeDateCollection := db.Database(DatabaseName).Collection(changeDateCollectionName)
	passportsCollection := db.Database(DatabaseName).Collection(passportsCollectionName)
	return &passportRepositoryImpl{
		client:               db,
		cache:                cache.New(cacheExpirationTime, cacheCleanUpTime),
		changeDateCollection: changeDateCollection,
		passportCollections:  passportsCollection,
	}
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
