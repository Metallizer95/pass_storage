package routers

import (
	domainpassport "store_server/internal/domain/passport"
	"store_server/internal/domain/routers"
	"store_server/internal/usecase/passport"
)

func RouteToRoutePassportsModel(route routers.ViksRoute) RoutePassportsModel {
	var routePassport RoutePassportsModel
	passportMapper := passport.NewMapper()
	routePassport.ViksRouteID = route.ViksRoutedID
	for _, pass := range route.SectionSet {
		p := passportMapper.ToPassportModel(pass)
		routePassport.Passports = append(routePassport.Passports, *p)
	}
	return routePassport
}

func RouteToModel(route routers.ViksRoute) RouteModel {
	var sectionSet []SectionModel
	for _, section := range route.SectionSet {
		s := SectionModel{
			SiteId:      section.SiteID,
			Sequence:    section.Sequence,
			SectionId:   section.SectionID,
			SectionName: section.SectionName,
			ChangeData:  section.ChangeDate,
			WorkType:    section.WorkType,
		}
		sectionSet = append(sectionSet, s)
	}
	model := RouteModel{
		MasterPmNum:    route.MasterPMNum,
		TripChangeData: route.TripChangeData,
		TripType:       route.TripType,
		ViksRouteID:    route.ViksRoutedID,
		Car:            route.Car,
		CarID:          route.CarID,
		Description:    route.Description,
		Eigthnum:       route.Eigthnum,
		SectionSetModel: SectionSetModel{
			Section: sectionSet,
		},
	}
	return model
}

func ModelToRoute(model RouteModel, passports []domainpassport.Passport) routers.ViksRoute {
	entity := routers.ViksRoute{
		MasterPMNum:    model.MasterPmNum,
		TripChangeData: model.TripChangeData,
		TripType:       model.TripType,
		ViksRoutedID:   model.ViksRouteID,
		Car:            model.Car,
		CarID:          model.CarID,
		Description:    model.Description,
		Eigthnum:       model.Eigthnum,
		SectionSet:     passports,
	}
	return entity
}

func ListRouteToModel(routes []routers.ViksRoute) ListRoutesModel {
	result := ListRoutesModel{}

	for _, route := range routes {
		routeModel := RouteToModel(route)
		result.RouteModel = append(result.RouteModel, InformationRouteModel{
			MasterPmNum:    routeModel.MasterPmNum,
			TripChangeData: routeModel.TripChangeData,
			TripType:       routeModel.TripType,
			ViksRouteID:    routeModel.ViksRouteID,
			Car:            routeModel.Car,
			CarID:          routeModel.CarID,
			Description:    routeModel.Description,
			Eigthnum:       routeModel.Eigthnum,
		})
	}
	return result
}
