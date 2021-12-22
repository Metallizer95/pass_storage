package common

type (
	UseCases interface {
		LoadPassportsByRouteUseCase() LoadPassportsByRouteUseCase
	}

	LoadPassportsByRouteUseCase interface {
		Load(routeid string) *Model
	}
)

//TODO: Remove link of Passports and Routes
