package routers

import "encoding/xml"

type RouteModel struct {
	XMLName        xml.Name   `xml:"VIKSROUTE"`
	Text           string     `xml:",chardata"`
	MasterPmNum    string     `xml:"MASTERPMNUM,attr"`
	TripChangeData string     `xml:"TripCHANGEDATA,attr"`
	ViksRouteID    string     `xml:"VIKSROUTEID,attr"`
	Car            string     `xml:"car,attr"`
	CarID          string     `xml:"carID,attr"`
	Description    string     `xml:"description,attr"`
	Eigthnum       string     `xml:"eightnum,attr"`
	SectionSet     SectionSet `xml:"SectionSet"`
}

type SectionSet struct {
	Text    string    `xml:",chardata"`
	Section []Section `xml:"Section"`
}

type Section struct {
	Text        string `xml:",chardata"`
	SiteId      string `xml:"siteId"`
	Sequence    string `xml:"SEQUENCE"`
	SectionId   string `xml:"sectionId"`
	SectionName string `xml:"sectionName"`
	ChangeData  string `xml:"CHANGEDATA"`
	WorkType    string `xml:"WorkType"`
}
