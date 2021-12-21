package routers

import (
	"store_server/internal/domain/routers"
)

type useCasesImpl struct {
	saveRouter     SaveRouterUseCase
	loadRouterByID LoadRouterByIDUseCase
	loadRouters    LoadRoutersUseCase
}

func NewUseCases(repository routers.Repository) UseCases {
	mng := routers.NewRouteManager(repository)
	saveUC := NewSaveRouterUseCaseImpl(mng)
	loadRouterByIdUC := NewLoadRouterByID(mng)
	loadRoutersUC := NewLoadRoutersUseCase(mng)
	return &useCasesImpl{
		saveRouter:     saveUC,
		loadRouterByID: loadRouterByIdUC,
		loadRouters:    loadRoutersUC,
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

// SaveRouter implementation
type saveRouterUseCaseImpl struct {
	mng routers.Manager
}

func (s *saveRouterUseCaseImpl) Save(route RouteModel) *RouteModel {
	entity := ModelToRoute(route)
	response := s.mng.SaveRoute(entity)
	if response == nil {
		return nil
	}
	responseModel := RouteToModel(*response)
	return &responseModel
}

func NewSaveRouterUseCaseImpl(mng routers.Manager) SaveRouterUseCase {
	return &saveRouterUseCaseImpl{mng: mng}
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
