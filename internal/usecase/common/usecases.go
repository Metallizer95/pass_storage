package common

import (
	"store_server/internal/domain/common"
)

// Use case implementation

type useCaseImpl struct {
	loadPassportsByRoute LoadPassportsByRouteUseCase
}

func NewUseCases(repository common.Repository) UseCases {
	mng := common.NewManager(repository)
	loadPassportsUC := newLoadPassportsByRouteUseCase(mng)
	return &useCaseImpl{loadPassportsByRoute: loadPassportsUC}
}

func (uc *useCaseImpl) LoadPassportsByRouteUseCase() LoadPassportsByRouteUseCase {
	return uc.loadPassportsByRoute
}

// LoadPassportsByRouteUseCase implementation

type loadPassportsByRouteUseCaseImpl struct {
	mng common.Manager
}

func newLoadPassportsByRouteUseCase(mng common.Manager) LoadPassportsByRouteUseCase {
	return &loadPassportsByRouteUseCaseImpl{mng: mng}
}

func (l *loadPassportsByRouteUseCaseImpl) Load(routeid string) *Model {
	passports := l.mng.FindPassports(routeid)
	passModel := EntityToModel(passports)
	return &passModel
}
