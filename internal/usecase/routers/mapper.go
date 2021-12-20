package routers

import (
	"store_server/internal/domain/routers"
)

func RouteToModel(route routers.ViksRoute) RouteModel {
	var sectionSet []SectionModel
	for _, section := range route.SectionSet {
		s := SectionModel{
			SiteId:      section.SiteId,
			Sequence:    section.Sequence,
			SectionId:   section.SectionId,
			SectionName: section.SectionName,
			ChangeData:  section.ChangeData,
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

func ModelToRoute(model RouteModel) routers.ViksRoute {
	var sectionSet []routers.Section
	for _, section := range model.SectionSetModel.Section {
		s := routers.Section{
			SiteId:      section.SiteId,
			Sequence:    section.Sequence,
			SectionId:   section.SectionId,
			SectionName: section.SectionName,
			ChangeData:  section.ChangeData,
			WorkType:    section.WorkType,
		}
		sectionSet = append(sectionSet, s)
	}
	entity := routers.ViksRoute{
		MasterPMNum:    model.MasterPmNum,
		TripChangeData: model.TripChangeData,
		TripType:       model.TripType,
		ViksRoutedID:   model.ViksRouteID,
		Car:            model.Car,
		CarID:          model.CarID,
		Description:    model.Description,
		Eigthnum:       model.Eigthnum,
		SectionSet:     sectionSet,
	}
	return entity
}

func ListRouteToModel(routes []*routers.ViksRoute) RoutesModel {
	result := RoutesModel{}

	for _, route := range routes {
		result.RouteModel = append(result.RouteModel, RouteToModel(*route))
	}
	return result
}
