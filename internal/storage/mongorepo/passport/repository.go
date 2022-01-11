package passportrepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"store_server/internal/domain/passport"
	"store_server/internal/storage/mongorepo"
	"store_server/pkg/cache"
	"store_server/pkg/logging"
	"strings"
	"time"
)

type PassportRepository interface {
	Create(passport passport.Data) *passport.Passport
	Read(id string) *passport.Passport
	ReadAll() []passport.Passport
	Update(passport passport.Passport) *passport.Passport
	Delete(id string) *passport.Passport
}

type passportRepositoryImpl struct {
	client               *mongo.Client
	cache                Cache
	logger               *logging.Logger
	changeDateCollection *mongo.Collection
	passportCollections  *mongo.Collection
}

func NewPassportRepository(db *mongo.Client) PassportRepository {
	cacheExpirationTime := 10 * time.Minute
	cacheCleanUpTime := 10 * time.Minute
	changeDateCollection := db.Database(mongorepo.DatabaseName).Collection(mongorepo.ChangeDateCollectionName)
	passportsCollection := db.Database(mongorepo.DatabaseName).Collection(mongorepo.PassportsCollectionName)

	logger, err := logging.GetLogger()
	if err != nil {
		panic(fmt.Sprintf("logger is unavailable: %v", err))
	}

	return &passportRepositoryImpl{
		client:               db,
		cache:                cache.New(cacheExpirationTime, cacheCleanUpTime),
		logger:               logger,
		changeDateCollection: changeDateCollection,
		passportCollections:  passportsCollection,
	}
}

func (m *passportRepositoryImpl) Create(d passport.Data) *passport.Passport {
	p := passport.Passport{
		ID:   d.SectionID,
		Data: d,
	}
	passportModel := passportToModel(p)

	_, err := m.passportCollections.InsertOne(context.TODO(), passportModel)
	if err != nil && !strings.Contains(err.Error(), "duplicate") {
		m.logger.Error(err)
		return nil
	}
	return &p
}

func (m *passportRepositoryImpl) Read(id string) *passport.Passport {
	p, ok := m.findByIdPassport(id)
	if !ok {
		return nil
	}

	result := modelToPassport(*p)
	return &result
}

func (m *passportRepositoryImpl) ReadAll() []passport.Passport {
	passports := m.findAllPassports()
	var result []passport.Passport

	for _, p := range passports {
		result = append(result, modelToPassport(p))
	}
	return nil
}

func (m *passportRepositoryImpl) Update(p passport.Passport) *passport.Passport {
	passportModel := passportToModel(p)
	if err := m.updatePassport(passportModel); err != nil {
		return nil
	}
	return &p
}

func (m *passportRepositoryImpl) Delete(id string) *passport.Passport {
	result, err := m.deletePassport(id)
	if err != nil {
		m.logger.Errorf("delete passport error: %v", err)
		return nil
	}
	p := modelToPassport(*result)
	return &p
}
