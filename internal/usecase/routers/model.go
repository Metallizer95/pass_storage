package routers

import (
	"encoding/xml"
	"store_server/internal/usecase/passport"
)

type ListRoutesModel struct {
	XMLName    xml.Name                `xml:"VIKSROUTERS"`
	RouteModel []InformationRouteModel `xml:"Routes"`
}

type InformationRouteModel struct {
	XMLName        xml.Name `xml:"VIKSROUTE"`
	Text           string   `xml:",chardata"`
	MasterPmNum    string   `xml:"MASTERPMNUM,attr"`
	TripChangeData string   `xml:"TripCHANGEDATA,attr"`
	TripType       string   `xml:"TripType,attr"`
	ViksRouteID    string   `xml:"VIKSROUTEID,attr"`
	Car            string   `xml:"car,attr"`
	CarID          string   `xml:"carID,attr"`
	Description    string   `xml:"description,attr"`
	Eigthnum       string   `xml:"eightnum,attr"`
}

type RouteModel struct {
	XMLName         xml.Name        `xml:"VIKSROUTE"`
	Text            string          `xml:",chardata"`
	MasterPmNum     string          `xml:"MASTERPMNUM,attr"`
	TripChangeData  string          `xml:"TripCHANGEDATA,attr"`
	TripType        string          `xml:"TripType,attr"`
	ViksRouteID     string          `xml:"VIKSROUTEID,attr"`
	Car             string          `xml:"car,attr"`
	CarID           string          `xml:"carID,attr"`
	Description     string          `xml:"description,attr"`
	Eigthnum        string          `xml:"eightnum,attr"`
	SectionSetModel SectionSetModel `xml:"SectionSet"`
}

type SectionSetModel struct {
	Text    string         `xml:",chardata"`
	Section []SectionModel `xml:"Section"`
}

type SectionModel struct {
	Text        string `xml:",chardata"`
	SiteId      string `xml:"siteId"`
	Sequence    string `xml:"SEQUENCE"`
	SectionId   string `xml:"sectionId"`
	SectionName string `xml:"sectionName"`
	ChangeData  string `xml:"CHANGEDATE"`
	WorkType    string `xml:"WorkType"`
}

type RoutePassportsModel struct {
	XMLName     xml.Name         `xml:"RoutePassports"`
	ViksRouteID string           `xml:"VIKSROUTEID"`
	Passports   []passport.Model `xml:"Passports"`
}
