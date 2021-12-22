package routers

import "store_server/internal/domain/passport"

type ViksRoute struct {
	MasterPMNum    string
	TripChangeData string
	TripType       string
	ViksRoutedID   string
	Car            string
	CarID          string
	Description    string
	Eigthnum       string
	SectionSet     []passport.Passport
}
