package common

import (
	"store_server/internal/domain/common"
	"store_server/internal/usecase/passport"
)

func EntityToModel(pr common.PassportsRoute) Model {
	var passports []passport.Model
	passportMapper := passport.NewMapper()
	for _, p := range pr.Passports {
		passports = append(passports, *passportMapper.ToPassportModel(p))
	}

	return Model{
		ViksRouteID: pr.ViksRouteID,
		Passports:   passports,
	}
}
