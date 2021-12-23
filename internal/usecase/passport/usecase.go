package passport

import "store_server/internal/domain/passport"

type useCasesImpl struct {
	saveUseCase      SavePassportUseCase
	loadUseCase      LoadPassportUseCase
	getTowersUseCase GetTowersUseCase
}

func (u *useCasesImpl) SavePassportUseCase() SavePassportUseCase {
	return u.saveUseCase
}

func (u *useCasesImpl) LoadPassportUseCase() LoadPassportUseCase {
	return u.loadUseCase
}

func (u *useCasesImpl) GetTowersUseCase() GetTowersUseCase {
	return u.getTowersUseCase
}

func NewUseCases(mng passport.Manager) UseCases {
	m := mapper{}
	loadUC := newPassportLoadCase(mng, m)
	saveUC := newSavePassportUseCase(mng, m)
	towersUC := newGetTowerUseCaseImpl(mng, m)
	return &useCasesImpl{
		saveUseCase:      saveUC,
		loadUseCase:      loadUC,
		getTowersUseCase: towersUC,
	}
}

// Save passport implementation
type savePassportUseCaseImpl struct {
	mng    passport.Manager
	mapper mapper
}

func (s *savePassportUseCaseImpl) Save(passport Model) *Model {
	r := s.mapper.ToPassportModel(*s.mng.SavePassport(*s.mapper.ToPassportData(passport)))
	return r
}

func newSavePassportUseCase(mng passport.Manager, mapperImpl mapper) SavePassportUseCase {
	return &savePassportUseCaseImpl{
		mng:    mng,
		mapper: mapperImpl,
	}
}

// Load passport implementation
type loadPassportUseCaseImpl struct {
	mng    passport.Manager
	mapper mapper
}

func (l *loadPassportUseCaseImpl) Load(id string) *Model {
	result := l.mng.LoadPassportByID(id)
	if result == nil {
		return nil
	}

	return l.mapper.ToPassportModel(*result)
}

func newPassportLoadCase(mng passport.Manager, mapperImpl mapper) LoadPassportUseCase {
	return &loadPassportUseCaseImpl{
		mng:    mng,
		mapper: mapperImpl,
	}
}

type getTowersUseCaseImpl struct {
	mng    passport.Manager
	mapper mapper
}

func (g *getTowersUseCaseImpl) LoadTowers(id string) *TowersModel {
	p := g.mng.LoadPassportByID(id)
	if p == nil {
		return nil
	}
	tModel := g.mapper.ToTowersModel(*p)
	return &tModel
}

func newGetTowerUseCaseImpl(mng passport.Manager, m mapper) GetTowersUseCase {
	return &getTowersUseCaseImpl{mng: mng, mapper: m}
}
