package passportrepo

import "store_server/internal/domain/passport"

func passportToModel(p passport.Passport) RepositoryModel {
	var towersModel []TowerRepoModel
	for _, t := range p.Towers.Towers {
		towersModel = append(towersModel, towerToModel(t))
	}
	header := HeaderRepoModel{
		SiteID:           p.SiteID,
		SectionName:      p.SectionName,
		SectionID:        p.SectionID,
		EchName:          p.EchName,
		EchkName:         p.EchkName,
		Location:         p.Location,
		WayAmount:        p.WayAmount,
		CurrentWay:       p.CurrentWay,
		CurrentWayID:     p.CurrentWayID,
		ChangeDate:       p.ChangeDate,
		InitialMeter:     p.InitialMeter,
		InitialKm:        p.InitialKm,
		InitialPk:        p.InitialPk,
		InitialM:         p.InitialM,
		PlotLength:       p.PlotLength,
		SuspensionAmount: p.SuspensionAmount,
		WorkType:         p.WorkType,
		Sequence:         p.Sequence,
	}
	return RepositoryModel{
		ID:     header.SectionID,
		Header: header,
		Towers: towersModel,
	}
}

func modelToPassport(model RepositoryModel) passport.Passport {
	var towers passport.Towers
	towers.Towers = make(map[string]passport.Tower)
	for _, t := range model.Towers {
		towerModel := modelToTower(t)
		towers.Towers[towerModel.ID] = towerModel
	}
	header := passport.Header{
		SiteID:           model.Header.SiteID,
		SectionName:      model.Header.SectionName,
		SectionID:        model.Header.SectionID,
		EchName:          model.Header.EchName,
		EchkName:         model.Header.EchkName,
		Location:         model.Header.Location,
		WayAmount:        model.Header.WayAmount,
		CurrentWay:       model.Header.CurrentWay,
		CurrentWayID:     model.Header.CurrentWayID,
		ChangeDate:       model.Header.ChangeDate,
		InitialMeter:     model.Header.InitialMeter,
		InitialKm:        model.Header.InitialKm,
		InitialPk:        model.Header.InitialPk,
		InitialM:         model.Header.InitialM,
		PlotLength:       model.Header.PlotLength,
		SuspensionAmount: model.Header.SuspensionAmount,
		WorkType:         model.Header.WorkType,
		Sequence:         model.Header.Sequence,
	}
	return passport.Passport{
		ID: model.ID,
		Data: passport.Data{
			Header: header,
			Towers: towers,
		},
	}
}

func towerToModel(tower passport.Tower) TowerRepoModel {
	return TowerRepoModel{
		ID:             tower.ID,
		AssetNum:       tower.AssetNum,
		StopSeq:        tower.StopSeq,
		Km:             tower.Km,
		Pk:             tower.Pk,
		M:              tower.M,
		Type:           tower.Type,
		Turn:           tower.Turn,
		Radius:         tower.Radius,
		Number:         tower.Number,
		Distance:       tower.Distance,
		Zigzag:         tower.Zigzag,
		Height:         tower.Height,
		Offset:         tower.Offset,
		Grounded:       tower.Grounded,
		Speed:          tower.Speed,
		SuspensionType: tower.SuspensionType,
		Catenary:       tower.Catenary,
		WireType:       tower.WireType,
		CountWire:      tower.CountWire,
		Longitude:      tower.Longitude,
		Latitude:       tower.Latitude,
		Gabarit:        tower.Gabarit,
	}
}

func modelToTower(model TowerRepoModel) passport.Tower {
	return passport.Tower{
		ID:             model.ID,
		AssetNum:       model.AssetNum,
		StopSeq:        model.StopSeq,
		Km:             model.Km,
		Pk:             model.Pk,
		M:              model.M,
		Type:           model.Type,
		Turn:           model.Turn,
		Radius:         model.Radius,
		Number:         model.Number,
		Distance:       model.Distance,
		Zigzag:         model.Zigzag,
		Height:         model.Height,
		Offset:         model.Offset,
		Grounded:       model.Grounded,
		Speed:          model.Speed,
		SuspensionType: model.SuspensionType,
		Catenary:       model.Catenary,
		WireType:       model.WireType,
		CountWire:      model.CountWire,
		Longitude:      model.Longitude,
		Latitude:       model.Latitude,
		Gabarit:        model.Gabarit,
	}
}

func passportToChangeDateModel(p passport.Passport) ChangeDateCollectionModel {
	return ChangeDateCollectionModel{
		PassportID: p.ID,
		ChangeDate: p.ChangeDate,
	}
}
