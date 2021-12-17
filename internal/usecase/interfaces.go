package usecase

import "store_server/internal/usecase/passport"

type (
	LoadPassportUseCase interface {
		Load(id int) (passport.PassportModel, error)
	}
)
