package routers

import (
	"store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
)

type useCasesImpl struct {
	saveRouter           SaveRouterUseCase
	loadRouterByID       LoadRouterByIDUseCase
	loadRouters          LoadRoutersUseCase
	loadPassportsByRoute LoadPassportsByRouteUseCase
}

func NewUseCases(repository routers.Repository) UseCases {
	mng := routers.NewRouteManager(repository)
	saveUC := NewSaveRouterUseCaseImpl(mng)
	loadRouterByIdUC := NewLoadRouterByID(mng)
	loadRoutersUC := NewLoadRoutersUseCase(mng)
	loadPassportsUC := NewLoadPassportByRouteUseCase(mng)
	return &useCasesImpl{
		saveRouter:           saveUC,
		loadRouterByID:       loadRouterByIdUC,
		loadRouters:          loadRoutersUC,
		loadPassportsByRoute: loadPassportsUC,
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

func (uc *useCasesImpl) LoadPassportsByRoute() LoadPassportsByRouteUseCase {
	return uc.loadPassportsByRoute
}

// SaveRouter implementation
type saveRouterUseCaseImpl struct {
	mng routers.Manager
}

func (s *saveRouterUseCaseImpl) Save(route RouteModel) *RouteModel {
	return nil
}

func NewSaveRouterUseCaseImpl(mng routers.Manager) SaveRouterUseCase {
	return &saveRouterUseCaseImpl{mng: mng}
}

// LoadRouterByID implementation
type loadRouterByIDUseCaseImpl struct {
	mng routers.Manager
}

func (l *loadRouterByIDUseCaseImpl) Load(id string) *RouteModel {
	return nil
}

func NewLoadRouterByID(mng routers.Manager) LoadRouterByIDUseCase {
	return &loadRouterByIDUseCaseImpl{mng: mng}
}

// LoadRouters implementation
type loadRoutersUseCaseImpl struct {
	mng routers.Manager
}

func (l *loadRoutersUseCaseImpl) Load() *RoutesModel {
	return nil
}

func NewLoadRoutersUseCase(mng routers.Manager) LoadRoutersUseCase {
	return &loadRoutersUseCaseImpl{mng: mng}
}

// LoadPassports implementation
type loadPassportsByRouteUseCaseImpl struct {
	mng routers.Manager
}

func (l *loadPassportsByRouteUseCaseImpl) Load(model RouteModel) []*passport.Passport {
	return nil
}

func NewLoadPassportByRouteUseCase(mng routers.Manager) LoadPassportsByRouteUseCase {
	return &loadPassportsByRouteUseCaseImpl{mng: mng}
}
