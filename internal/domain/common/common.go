package common

import (
	"store_server/internal/domain/passport"
)

type PassportsRoute struct {
	ViksRouteID string
	Passports   []passport.Passport
}

// TODO: Remove common entity
