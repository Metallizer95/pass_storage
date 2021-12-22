package passport

import "store_server/internal/domain/passport"

type useCasesImpl struct {
	saveUseCase SavePassportUseCase
	loadUseCase LoadPassportUseCase
}

func (u *useCasesImpl) SavePassport() SavePassportUseCase {
	return u.saveUseCase
}

func (u *useCasesImpl) LoadPassport() LoadPassportUseCase {
	return u.loadUseCase
}

func NewUseCases(repository passport.Repository) UseCases {
	m := mapper{}
	mng := passport.NewPassportManager(repository)
	loadUC := NewPassportLoadCase(mng, m)
	saveUC := NewSavePassportUseCase(mng, m)
	return &useCasesImpl{
		saveUseCase: saveUC,
		loadUseCase: loadUC,
	}
}

// Save passport implementation
type savePassportUseCaseImpl struct {
	mng passport.Manager
	m   mapper
}

func (s *savePassportUseCaseImpl) Save(passport Model) *Model {
	r := s.m.ToPassportModel(*s.mng.SavePassport(*s.m.ToPassportData(passport)))
	return r
}

func NewSavePassportUseCase(mng passport.Manager, mapperImpl mapper) SavePassportUseCase {
	return &savePassportUseCaseImpl{
		mng: mng,
		m:   mapperImpl,
	}
}

// Load passport implementation
type loadPassportUseCaseImpl struct {
	mng passport.Manager
	m   mapper
}

func (l *loadPassportUseCaseImpl) Load(id string) *Model {
	result := l.mng.LoadPassportByID(id)
	if result == nil {
		return nil
	}

	return l.m.ToPassportModel(*result)
}

func NewPassportLoadCase(mng passport.Manager, mapperImpl mapper) LoadPassportUseCase {
	return &loadPassportUseCaseImpl{
		mng: mng,
		m:   mapperImpl,
	}
}

type getTowersUseCaseImpl struct {
	mng passport.Manager
	m   mapper
}

func (g *getTowersUseCaseImpl) LoadTowers(id string) *Towers {
	return nil
}
