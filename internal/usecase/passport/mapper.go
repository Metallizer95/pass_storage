package passport

import (
	"store_server/internal/domain/passport"
)

type Mapper interface {
	ToPassport(p PassportModel) passport.Passport
	ToPassportModel(p passport.Passport) PassportModel
}
