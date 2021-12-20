package routers

import "store_server/internal/domain/passport"

type (
	UseCases interface {
		SaveRouter() SaveRouterUseCase
		LoadRouterByID() LoadRouterByIDUseCase
		LoadRouters() LoadRoutersUseCase
		LoadPassportsByRoute() LoadPassportsByRouteUseCase
	}

	SaveRouterUseCase interface {
		Save(route RouteModel) *RouteModel
	}

	LoadRouterByIDUseCase interface {
		Load(id string) *RouteModel
	}

	LoadRoutersUseCase interface {
		Load() *RoutesModel
	}

	LoadPassportsByRouteUseCase interface {
		Load(model RouteModel) []*passport.Passport
	}
)

//TODO: Загрузка маршрутов - метод POST формат XML
//TODO:  Получение списка маршрутов - метод  GET (формат необходимо соглсовать)
//TODO:  При выборе маршрута получение всех паспортов, которые принадлежат маршруту - метод GET  формат архив файлов
//TODO: Нахождение просроченных паспортов путём сравнение в каждом паспорте параметра change date с текущей датой, вернуть id просроченных паспортов, формат json
