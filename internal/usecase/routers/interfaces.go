package routers

type (
	UseCases interface {
		SaveRouter() SaveRouterUseCase
		LoadRouterByID() LoadRouterByIDUseCase
		LoadRouters() LoadRoutersUseCase
		LoadPassportsByRoute() LoadPassportsByRoute
	}

	SaveRouterUseCase interface {
		Save(route RouteModel) (*RouteModel, error)
	}

	LoadRouterByIDUseCase interface {
		Load(id string) *RouteModel
	}

	LoadRoutersUseCase interface {
		Load() *ListRoutesModel
	}

	LoadPassportsByRoute interface {
		Load(routeid string) *RoutePassportsModel
	}
)
