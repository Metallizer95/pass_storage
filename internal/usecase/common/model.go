package common

import (
	"encoding/xml"
	"store_server/internal/usecase/passport"
)

type Model struct {
	XMLName     xml.Name         `xml:"RoutePassports"`
	ViksRouteID string           `xml:"VIKSROUTEID"`
	Passports   []passport.Model `xml:"Passports"`
}
