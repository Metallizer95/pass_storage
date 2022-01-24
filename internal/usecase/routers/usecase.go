package routers

import (
	"fmt"
	domainpassport "store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
	"store_server/internal/usecase/errs"
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

// TODO: I don`t like this solution. Route entity consist passports instead just ID on them
func (s *saveRouterUseCaseImpl) Save(route RouteModel) (*RouteModel, error) {
	var passports []domainpassport.Passport
	passportMapper := passport.NewMapper()

	var noExistedPassports []string
	for _, pass := range route.SectionSetModel.Section {
		pModel := s.passportUseCases.LoadPassportUseCase().Load(pass.SectionId)

		if pModel == nil {
			noExistedPassports = append(noExistedPassports, pass.SectionId)
			continue
		}
		pModel.Header.WorkType = pass.WorkType
		pModel.Header.Sequence = pass.Sequence
		p := passportMapper.ToPassport(pModel.Model)
		passports = append(passports, *p)
	}
	if len(noExistedPassports) != 0 {
		return nil, fmt.Errorf("missing passports in database with id: %v", noExistedPassports)
	}

	entity := ModelToRoute(route, passports)
	response := s.mng.SaveRoute(entity)
	if response == nil {
		return nil, errs.ErrObjectAlreadyExists
	}
	responseModel := RouteToModel(*response)
	return &responseModel, nil
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

type loadPassportsByRouteImpl struct {
	mng routers.Manager
}

func newLoadPassportsByRoute(mng routers.Manager) LoadPassportsByRoute {
	return &loadPassportsByRouteImpl{mng}
}

func (l *loadPassportsByRouteImpl) Load(routeid string) *RoutePassportsModel {
	// TODO: Which layer should have responsibility for linking passport and route repos
	route := l.mng.LoadRouteByID(routeid)
	if route == nil {
		return nil
	}
	result := RouteToRoutePassportsModel(*route)
	return &result
}
