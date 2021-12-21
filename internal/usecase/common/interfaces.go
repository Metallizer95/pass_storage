package common

type (
	UseCases interface {
		LoadPassportsByRouteUseCase() LoadPassportsByRouteUseCase
	}

	LoadPassportsByRouteUseCase interface {
		Load(routeid string) *Model
	}
)
