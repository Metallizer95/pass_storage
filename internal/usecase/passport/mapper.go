package passport

import (
	"store_server/internal/domain/passport"
)

type Mapper interface {
	ToPassport(p Model) *passport.Data
	ToPassportModel(p passport.Passport) *Model
}

type mapper struct {
}

// TODO: how I can copy structures by less code

func (m *mapper) ToPassport(p Model) *passport.Data {
	var towers passport.Towers
	for _, t := range p.Towers.Towers {
		towers.Towers = append(towers.Towers, passport.Tower{
			ID:             t.ID,
			AssetNum:       t.AssetNum,
			StopSeq:        t.StopSeq,
			Km:             t.Km,
			Pk:             t.Pk,
			M:              t.M,
			Type:           t.Type,
			Turn:           t.Turn,
			Radius:         t.Radius,
			Number:         t.Number,
			Distance:       t.Distance,
			Zigzag:         t.Zigzag,
			Height:         t.Height,
			Offset:         t.Offset,
			Grounded:       t.Grounded,
			Speed:          t.Speed,
			SuspensionType: t.SuspensionType,
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
			Location:         p.Header.Location,
			WayAmount:        p.Header.WayAmount,
			CurrentWay:       p.Header.CurrentWay,
			CurrentWayID:     p.Header.CurrentWayID,
			ChangeData:       p.Header.ChangeData,
			InitialMeter:     p.Header.InitialMeter,
			InitialKm:        p.Header.InitialKm,
			InitialPk:        p.Header.InitialPk,
			InitialM:         p.Header.InitialM,
			PlotLength:       p.Header.PlotLength,
			SuspensionAmount: p.Header.SuspensionAmount,
		},
		Towers: towers,
	}
}

func (m *mapper) ToPassportModel(p passport.Passport) *Model {
	var towers Towers
	for _, t := range p.Towers.Towers {
		towers.Towers = append(towers.Towers, Tower{
			ID:             t.ID,
			AssetNum:       t.AssetNum,
			StopSeq:        t.StopSeq,
			Km:             t.Km,
			Pk:             t.Pk,
			M:              t.M,
			Type:           t.Type,
			Turn:           t.Turn,
			Radius:         t.Radius,
			Number:         t.Number,
			Distance:       t.Distance,
			Zigzag:         t.Zigzag,
			Height:         t.Height,
			Offset:         t.Offset,
			Grounded:       t.Grounded,
			Speed:          t.Speed,
			SuspensionType: t.SuspensionType,
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
		Header: Header{
			SiteID:           h.SiteID,
			SectionName:      h.SectionName,
			SectionID:        h.SectionID,
			EchName:          h.EchName,
			EchkName:         h.EchkName,
			Location:         h.Location,
			WayAmount:        h.WayAmount,
			CurrentWay:       h.CurrentWay,
			CurrentWayID:     h.CurrentWayID,
			ChangeData:       h.ChangeData,
			InitialMeter:     h.InitialMeter,
			InitialKm:        h.InitialKm,
			InitialPk:        h.InitialPk,
			InitialM:         h.InitialM,
			PlotLength:       h.PlotLength,
			SuspensionAmount: h.SuspensionAmount,
		},
		Towers: towers,
	}
}
