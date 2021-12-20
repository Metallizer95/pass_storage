package routers

import "encoding/xml"

// TODO: change response format to zip archive (load from db, squeeze and send to client)

type RoutesModel struct {
	XMLName    xml.Name     `xml:"VIKSROUTERS,omitempty"`
	RouteModel []RouteModel `xml:"Routes"`
}

type RouteModel struct {
	XMLName         xml.Name        `xml:"VIKSROUTE,omitempty"`
	Text            string          `xml:",chardata,omitempty"`
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
	Text    string         `xml:",chardata,omitempty"`
	Section []SectionModel `xml:"Section"`
}

type SectionModel struct {
	Text        string `xml:",chardata,omitempty"`
	SiteId      string `xml:"siteId"`
	Sequence    string `xml:"SEQUENCE"`
	SectionId   string `xml:"sectionId"`
	SectionName string `xml:"sectionName"`
	ChangeData  string `xml:"CHANGEDATA"`
	WorkType    string `xml:"WorkType"`
}
