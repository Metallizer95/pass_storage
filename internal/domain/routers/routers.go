package routers

type ViksRoute struct {
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

// TODO: Link passport and route by SectionSet. Change SectionSet to passport entity
