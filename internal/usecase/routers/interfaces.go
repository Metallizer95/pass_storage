package routers

type (
	UseCases interface {
		SaveRouter() SaveRouterUseCase
		LoadRouterByID() LoadRouterByIDUseCase
		LoadRouters() LoadRoutersUseCase
	}

	SaveRouterUseCase interface {
		Save(route RouteModel) *RouteModel
	}

	LoadRouterByIDUseCase interface {
		Load(id string) *RouteModel
	}

	LoadRoutersUseCase interface {
		Load() *ListRoutesModel
	}
)
