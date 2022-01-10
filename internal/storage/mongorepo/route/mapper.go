package routerepo

import (
	"store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
)

func routeToRepositoryModel(route routers.ViksRoute) repositoryModel {
	var sectionSet []string
	for _, s := range route.SectionSet {
		sectionSet = append(sectionSet, s.ID)
	}
	return repositoryModel{
		ID: route.ViksRoutedID,
		Route: routeRepositoryModel{
			MasterPMNum:    route.MasterPMNum,
			TripChangeDate: route.TripChangeData,
			TripType:       route.TripType,
			ViksRoutedID:   route.ViksRoutedID,
			Car:            route.Car,
			CarID:          route.CarID,
			Description:    route.Description,
			Eigthnum:       route.Eigthnum,
			SectionSet:     sectionSet,
		},
	}
}

func repositoryModelToRoute(model repositoryModel) routers.ViksRoute {
	var sectionSet []passport.Passport

	for _, s := range model.Route.SectionSet {
		sectionSet = append(sectionSet, passport.Passport{ID: s})
	}

	return routers.ViksRoute{
		MasterPMNum:    model.Route.MasterPMNum,
		TripChangeData: model.Route.TripChangeDate,
		TripType:       model.Route.TripType,
		ViksRoutedID:   model.Route.ViksRoutedID,
		Car:            model.Route.Car,
		CarID:          model.Route.CarID,
		Description:    model.Route.Description,
		Eigthnum:       model.Route.Eigthnum,
		SectionSet:     sectionSet,
	}
}
