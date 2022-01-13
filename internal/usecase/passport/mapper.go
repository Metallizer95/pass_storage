package passport

import (
	"store_server/internal/domain/passport"
)

type Mapper interface {
	ToPassportData(p Model) *passport.Data
	ToPassport(p Model) *passport.Passport
	ToPassportModel(p passport.Passport) *Model
	ToTowersModel(p passport.Towers, sectionId string) TowersModel
	ToTowerModel(tower passport.Tower) TowerModel
	PassportToExpiredModel(p passport.Passport, duration string) ExpiredModel
	ListExpiredModelToExpiredPassportModel(expireModels []ExpiredModel) ExpiredPassportsModel
}

type mapper struct {
}

func NewMapper() Mapper {
	return &mapper{}
}

func (m *mapper) ToPassportData(p Model) *passport.Data {
	var towers passport.Towers
	towers.Towers = make(map[string]passport.Tower)
	for _, t := range p.Towers.Towers {
		towers.Towers[t.Idtf] = passport.Tower{
			ID:             t.Idtf,
			AssetNum:       t.AssetNum,
			StopSeq:        t.StopSeq,
			Km:             t.Km,
			Pk:             t.Pk,
			M:              t.M,
			Type:           t.TFTYPE,
			Turn:           t.TURN,
			Radius:         t.RADIUS,
			Number:         t.Number,
			Distance:       t.Distance,
			Zigzag:         t.Zigzag,
			Height:         t.Height,
			Offset:         t.Offset,
			Grounded:       t.Grounded,
			Speed:          t.SPEED,
			SuspensionType: t.SuspensionType,
			Catenary:       t.Catenary,
			WireType:       t.WireType,
			CountWire:      t.CountWire,
			Longitude:      t.Longitude,
			Latitude:       t.Latitude,
			Gabarit:        t.Gabarit,
		}
	}
	return &passport.Data{
		Header: passport.Header{
			SiteID:           p.Header.SiteID,
			SectionName:      p.Header.SectionName,
			SectionID:        p.Header.SectionID,
			EchName:          p.Header.EchName,
			EchkName:         p.Header.EchkName,
			Location:         p.Header.LocationId,
			WayAmount:        p.Header.WayAmount,
			CurrentWay:       p.Header.CurrentWay,
			CurrentWayID:     p.Header.CurrentWayID,
			ChangeDate:       p.Header.CHANGEDATE,
			InitialMeter:     p.Header.InitialMeter,
			InitialKm:        p.Header.InitialKm,
			InitialPk:        p.Header.InitialPK,
			InitialM:         p.Header.InitialM,
			PlotLength:       p.Header.PlotLength,
			SuspensionAmount: p.Header.SuspensionAmount,
			Sequence:         p.Header.Sequence,
			WorkType:         p.Header.WorkType,
		},
		Towers: towers,
	}
}

func (m *mapper) ToPassportModel(p passport.Passport) *Model {
	var towers TowersModel
	for _, t := range p.Towers.Towers {
		towers.Towers = append(towers.Towers, TowerModel{
			Idtf:           t.ID,
			AssetNum:       t.AssetNum,
			StopSeq:        t.StopSeq,
			Km:             t.Km,
			Pk:             t.Pk,
			M:              t.M,
			TFTYPE:         t.Type,
			TURN:           t.Turn,
			RADIUS:         t.Radius,
			Number:         t.Number,
			Distance:       t.Distance,
			Zigzag:         t.Zigzag,
			Height:         t.Height,
			Offset:         t.Offset,
			Grounded:       t.Grounded,
			SPEED:          t.Speed,
			SuspensionType: t.SuspensionType, // TODO: change name of suspensionTime field
			Catenary:       t.Catenary,
			WireType:       t.WireType,
			CountWire:      t.CountWire,
			Longitude:      t.Longitude,
			Latitude:       t.Latitude,
			Gabarit:        t.Gabarit,
		})
	}
	h := p.Header
	return &Model{
		ID: p.ID,
		Header: HeaderModel{
			SiteID:           h.SiteID,
			SectionName:      h.SectionName,
			SectionID:        h.SectionID,
			EchName:          h.EchName,
			EchkName:         h.EchkName,
			LocationId:       h.Location,
			WayAmount:        h.WayAmount,
			CurrentWay:       h.CurrentWay,
			CurrentWayID:     h.CurrentWayID,
			CHANGEDATE:       h.ChangeDate,
			InitialMeter:     h.InitialMeter,
			InitialKm:        h.InitialKm,
			InitialPK:        h.InitialPk,
			InitialM:         h.InitialM,
			PlotLength:       h.PlotLength,
			SuspensionAmount: h.SuspensionAmount,
			Sequence:         h.Sequence,
			WorkType:         h.WorkType,
		},
		Towers: towers,
	}
}

func (m *mapper) ToPassport(p Model) *passport.Passport {
	if p.ID == "" {
		return nil
	}
	return &passport.Passport{
		ID:   p.ID,
		Data: *m.ToPassportData(p),
	}
}

func (m *mapper) ToTowersModel(p passport.Towers, sectionId string) TowersModel {
	var result TowersModel
	for _, t := range p.Towers {
		result.Towers = append(result.Towers, m.ToTowerModel(t))
	}
	result.SectionID = sectionId
	return result
}

func (m *mapper) ToTowerModel(tower passport.Tower) TowerModel {
	return TowerModel{
		Idtf:           tower.ID,
		AssetNum:       tower.AssetNum,
		StopSeq:        tower.StopSeq,
		Km:             tower.Km,
		Pk:             tower.Pk,
		M:              tower.M,
		TFTYPE:         tower.Type,
		TURN:           tower.Turn,
		RADIUS:         tower.Radius,
		Number:         tower.Number,
		Distance:       tower.Distance,
		Zigzag:         tower.Zigzag,
		Height:         tower.Height,
		Offset:         tower.Offset,
		Grounded:       tower.Grounded,
		SPEED:          tower.Speed,
		SuspensionType: tower.SuspensionType,
		Catenary:       tower.Catenary,
		WireType:       tower.WireType,
		CountWire:      tower.CountWire,
		Longitude:      tower.Longitude,
		Latitude:       tower.Latitude,
		Gabarit:        tower.Gabarit,
	}
}

func (m *mapper) PassportToExpiredModel(p passport.Passport, duration string) ExpiredModel {
	return ExpiredModel{
		ID:         p.SectionID,
		ChangeData: p.ChangeDate,
		Duration:   duration,
	}
}

func (m *mapper) ListExpiredModelToExpiredPassportModel(expireModels []ExpiredModel) ExpiredPassportsModel {
	var result ExpiredPassportsModel
	for _, e := range expireModels {
		result.Passports = append(result.Passports, e)
	}
	return result
}
