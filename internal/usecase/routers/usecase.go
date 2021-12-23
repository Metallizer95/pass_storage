package routers

import (
	domainpassport "store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
	"store_server/internal/usecase/passport"
)

type useCasesImpl struct {
	saveRouter           SaveRouterUseCase
	loadRouterByID       LoadRouterByIDUseCase
	loadRouters          LoadRoutersUseCase
	loadPassportsByRoute LoadPassportsByRoute
}

func NewUseCases(mng routers.Manager, passportUseCases passport.UseCases) UseCases {
	saveUC := NewSaveRouterUseCaseImpl(mng, passportUseCases)
	loadRouterByIdUC := NewLoadRouterByID(mng)
	loadRoutersUC := NewLoadRoutersUseCase(mng)
	loadPassportByRouteUC := newLoadPassportsByRoute(mng)
	return &useCasesImpl{
		saveRouter:           saveUC,
		loadRouterByID:       loadRouterByIdUC,
		loadRouters:          loadRoutersUC,
		loadPassportsByRoute: loadPassportByRouteUC,
	}
}

func (uc *useCasesImpl) SaveRouter() SaveRouterUseCase {
	return uc.saveRouter
}

func (uc *useCasesImpl) LoadRouterByID() LoadRouterByIDUseCase {
	return uc.loadRouterByID
}

func (uc *useCasesImpl) LoadRouters() LoadRoutersUseCase {
	return uc.loadRouters
}

func (uc *useCasesImpl) LoadPassportsByRoute() LoadPassportsByRoute {
	return uc.loadPassportsByRoute
}

// SaveRouter implementation
type saveRouterUseCaseImpl struct {
	mng              routers.Manager
	passportUseCases passport.UseCases
}

func (s *saveRouterUseCaseImpl) Save(route RouteModel) *RouteModel {
	var passports []domainpassport.Passport
	passportMapper := passport.NewMapper()
	for _, pass := range route.SectionSetModel.Section {
		pModel := s.passportUseCases.LoadPassportUseCase().Load(pass.SectionId)

		if pModel == nil {
			continue
		}
		pModel.Header.WorkType = pass.WorkType
		pModel.Header.Sequence = pass.Sequence
		p := passportMapper.ToPassport(*pModel)
		passports = append(passports, *p)
	}
	entity := ModelToRoute(route, passports)
	response := s.mng.SaveRoute(entity)
	if response == nil {
		return nil
	}
	responseModel := RouteToModel(*response)
	return &responseModel
}

func NewSaveRouterUseCaseImpl(mng routers.Manager, passportUseCase passport.UseCases) SaveRouterUseCase {
	return &saveRouterUseCaseImpl{mng: mng, passportUseCases: passportUseCase}
}

// LoadRouterByID implementation
type loadRouterByIDUseCaseImpl struct {
	mng routers.Manager
}

func (l *loadRouterByIDUseCaseImpl) Load(id string) *RouteModel {
	route := l.mng.LoadRouteByID(id)
	if route == nil {
		return nil
	}
	model := RouteToModel(*route)
	return &model
}

func NewLoadRouterByID(mng routers.Manager) LoadRouterByIDUseCase {
	return &loadRouterByIDUseCaseImpl{mng: mng}
}

// LoadRouters implementation
type loadRoutersUseCaseImpl struct {
	mng routers.Manager
}

func (l *loadRoutersUseCaseImpl) Load() *ListRoutesModel {
	routes := l.mng.LoadRoutes()
	if routes == nil {
		return nil
	}
	routesModel := ListRouteToModel(routes)
	return &routesModel
}

func NewLoadRoutersUseCase(mng routers.Manager) LoadRoutersUseCase {
	return &loadRoutersUseCaseImpl{mng: mng}
}

// LoadPassportsByRoute implementation

type loadPassportByRouteImpl struct {
	mng routers.Manager
}

func newLoadPassportsByRoute(mng routers.Manager) LoadPassportsByRoute {
	return &loadPassportByRouteImpl{mng}
}

func (l *loadPassportByRouteImpl) Load(routeid string) *RoutePassportsModel {
	route := l.mng.LoadRouteByID(routeid)
	if route == nil {
		return nil
	}
	result := RouteToRoutePassportsModel(*route)
	return &result
}
