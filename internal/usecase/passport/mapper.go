package passport

import (
	"store_server/internal/domain/passport"
)

type Mapper interface {
	ToPassportData(p Model) *passport.Data
	ToPassport(p Model) *passport.Passport
	ToPassportModel(p passport.Passport) *Model
}

type mapper struct {
}

func NewMapper() Mapper {
	return &mapper{}
}

// TODO: how I can copy structures by less code

type TowerMapper interface {
	ToTower(p Tower) *passport.Tower
	ToTowerModel(p passport.Tower) *Tower
}

func (m *mapper) ToPassportData(p Model) *passport.Data {
	var towers passport.Towers
	for _, t := range p.Towers.Towers {
		towers.Towers = append(towers.Towers, passport.Tower{
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
			SuspensionType: t.SuspensionType, //TODO: change time to type
			Catenary:       t.Catenary,
			WireType:       t.WireType,
			CountWire:      t.CountWire,
			Longitude:      t.Longitude,
			Latitude:       t.Latitude,
			Gabarit:        t.Gabarit,
		})
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
	var towers Towers
	for _, t := range p.Towers.Towers {
		towers.Towers = append(towers.Towers, Tower{
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
		Header: Header{
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
