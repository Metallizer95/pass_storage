package routers

import "encoding/xml"

type ViksRoute struct {
	XMLName        xml.Name
	Text           string
	MasterPMNum    string
	TripChangeData string
	TripType       string
	ViksRoutedID   string
	Car            string
	CarID          string
	Description    string
	Eigthnum       string
	SectionSet     []Section
}

type Section struct {
	SiteId      string
	Sequence    string
	SectionId   string
	SectionName string
	ChangeData  string
	WorkType    string
}
