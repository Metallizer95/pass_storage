package commonstore

import (
	"store_server/internal/domain/common"
	"store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
)

type repositoryCommonImpl struct {
	PassportRepo passport.Repository
	RouteRepo    routers.Repository
}

func NewRepository(passportRepo passport.Repository, routeRepo routers.Repository) common.Repository {
	return &repositoryCommonImpl{PassportRepo: passportRepo, RouteRepo: routeRepo}
}

func (r *repositoryCommonImpl) RouteRepository() routers.Repository {
	return r.RouteRepo
}

func (r *repositoryCommonImpl) PassportRepository() passport.Repository {
	return r.PassportRepo
}

func (r *repositoryCommonImpl) FindPassportsByRoute(routeid string) common.PassportsRoute {
	route := r.RouteRepo.Read(routeid)
	var result common.PassportsRoute

	for _, section := range route.SectionSet {
		pass := r.PassportRepo.Read(section.SectionId)
		if pass != nil {
			result.Passports = append(result.Passports, *pass)
		}
	}
	result.ViksRouteID = routeid
	return result
}
