package passport

import "store_server/internal/domain/passport"

type SavePassportUseCase interface {
	Save(passport Model) *passport.Passport
}

type SavePassportUseCaseImpl struct {
	mng passport.Manager
	m   mapper
}

func (s *SavePassportUseCaseImpl) Save(passport Model) *passport.Passport {
	return s.mng.SavePassport(*s.m.ToPassport(passport))
}

type LoadPassportUseCase interface {
	Load(id string) *Model
}

type LoadPassportUseCaseImpl struct {
	mng passport.Manager
	m   mapper
}

func (l *LoadPassportUseCaseImpl) Load(id string) *Model {
	result := l.mng.LoadPassportByID(id)
	if result == nil {
		return nil
	}

	return l.m.ToPassportModel(*result)
}
