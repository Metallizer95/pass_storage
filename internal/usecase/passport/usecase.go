package passport

import "store_server/internal/domain/passport"

type useCasesImpl struct {
	saveUseCase                       SavePassportUseCase
	loadUseCase                       LoadPassportUseCase
	getTowersUseCase                  GetTowersUseCase
	findTowerByIdAndCoordinateUseCase FindTowerByIdAndCoordinateUseCase
	findExpiredPassportsUseCase       FindExpiredPassportsUseCase
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

func (u *useCasesImpl) FindTowerByIdAndCoordinateUseCase() FindTowerByIdAndCoordinateUseCase {
	return u.findTowerByIdAndCoordinateUseCase
}

func (u *useCasesImpl) FindExpiredPassportsUseCase() FindExpiredPassportsUseCase {
	return u.findExpiredPassportsUseCase
}

func NewUseCases(mng passport.Manager) UseCases {
	m := &mapper{}
	loadUC := newPassportLoadCase(mng, m)
	saveUC := newSavePassportUseCase(mng, m)
	towersUC := newGetTowerUseCaseImpl(mng, m)
	towerByIdAndCoord := newFindTowerByIdAndCoordinateUseCase(mng, m)
	return &useCasesImpl{
		saveUseCase:                       saveUC,
		loadUseCase:                       loadUC,
		getTowersUseCase:                  towersUC,
		findTowerByIdAndCoordinateUseCase: towerByIdAndCoord,
	}
}

// Save passport implementation
type savePassportUseCaseImpl struct {
	mng    passport.Manager
	mapper Mapper
}

func (s *savePassportUseCaseImpl) Save(passport Model) *Model {
	r := s.mapper.ToPassportModel(*s.mng.SavePassport(*s.mapper.ToPassportData(passport)))
	return r
}

func newSavePassportUseCase(mng passport.Manager, mapperImpl Mapper) SavePassportUseCase {
	return &savePassportUseCaseImpl{
		mng:    mng,
		mapper: mapperImpl,
	}
}

// Load passport implementation
type loadPassportUseCaseImpl struct {
	mng    passport.Manager
	mapper Mapper
}

func (l *loadPassportUseCaseImpl) Load(id string) *Model {
	result := l.mng.LoadPassportByID(id)
	if result == nil {
		return nil
	}
	passportModel := l.mapper.ToPassportModel(*result)
	updateExpiration(passportModel)
	return passportModel
}

func newPassportLoadCase(mng passport.Manager, mapperImpl Mapper) LoadPassportUseCase {
	return &loadPassportUseCaseImpl{
		mng:    mng,
		mapper: mapperImpl,
	}
}

// Get all towers of passport implementation
type getTowersUseCaseImpl struct {
	mng    passport.Manager
	mapper Mapper
}

func (g *getTowersUseCaseImpl) LoadAllTowerByPassportId(id string) *TowersModel {
	p := g.mng.LoadPassportByID(id)
	if p == nil {
		return nil
	}
	towers := p.GetAllTowers()

	tModel := g.mapper.ToTowersModel(towers, p.SectionID)
	return &tModel
}

func (g *getTowersUseCaseImpl) LoadTowerById(passportId, towerId string) *TowerModel {
	p := g.mng.LoadPassportByID(passportId)
	if p == nil {
		return nil
	}

	tower := p.GetTowerById(towerId)
	if tower == nil {
		return nil
	}

	towerModel := g.mapper.ToTowerModel(*tower)
	return &towerModel
}

func newGetTowerUseCaseImpl(mng passport.Manager, m Mapper) GetTowersUseCase {
	return &getTowersUseCaseImpl{mng: mng, mapper: m}
}

// Find tower by id and coordinate implementation
type findTowerByIdAndCoordinateUseCaseImpl struct {
	mng    passport.Manager
	mapper Mapper
}

func (f findTowerByIdAndCoordinateUseCaseImpl) FindTower(id string, longitude float64, latitude float64) *TowerModel {
	p := f.mng.LoadPassportByID(id)
	if p == nil {
		return nil
	}
	tower := p.GetTowerByCoordinate(longitude, latitude)
	towerModel := f.mapper.ToTowerModel(*tower)
	return &towerModel
}

func newFindTowerByIdAndCoordinateUseCase(mng passport.Manager, m Mapper) FindTowerByIdAndCoordinateUseCase {
	return findTowerByIdAndCoordinateUseCaseImpl{mng, m}
}

// find expired passports implementation
type findExpiredPassportsUseCaseImpl struct {
	mng    passport.Manager
	mapper Mapper
}

// TODO: I think it is bad method in incorrect place. Need help of Kot

func (f *findExpiredPassportsUseCaseImpl) FindPassports() ExpiredPassportsModel {
	passports := f.mng.LoadAll()
	var result []ExpiredModel
	for _, p := range passports {
		if expired, duration := IsExpired(p.ChangeDate); expired {
			result = append(result, f.mapper.PassportToExpiredModel(p, duration))
		}
	}
	return f.mapper.ListExpiredModelToExpiredPassportModel(result)
}

func newFindExpiredPassports(mng passport.Manager, m Mapper) FindExpiredPassportsUseCase {
	return &findExpiredPassportsUseCaseImpl{mng, m}
}
