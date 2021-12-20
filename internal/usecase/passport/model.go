package passport

import (
	"encoding/xml"
)

type Router struct {
	Passports []Model `xml:"passports"`
}

//TODO: Which format does change data have?
//TODO: What is M?

type Model struct {
	XMLName xml.Name `xml:"Data"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,omitempty"`
	Header  Header   `xml:"Header"`
	Towers  Towers   `xml:"Towers"`
}

type Header struct {
	Text             string `xml:",chardata"`
	SiteID           int    `xml:"siteID"`
	SectionName      string `xml:"sectionName"`
	SectionID        int    `xml:"sectionID"`
	EchName          string `xml:"echName"`
	EchkName         string `xml:"echkName"`
	LocationId       string `xml:"locationId"`
	WayAmount        int    `xml:"wayAmount"`
	CurrentWay       int    `xml:"currentWay"`
	CurrentWayID     int    `xml:"currentWayID"`
	CHANGEDATA       string `xml:"CHANGEDATA"`
	InitialMeter     int    `xml:"initialMeter"`
	InitialKm        int    `xml:"initialKm"`
	InitialPK        int    `xml:"initialPK"`
	InitialM         int    `xml:"initialM"`
	PlotLength       int    `xml:"plotLength"`
	SuspensionAmount int    `xml:"suspensionAmount"`
}

type Towers struct {
	Text   string  `xml:",chardata"`
	Towers []Tower `xml:"Tower"`
}

type Tower struct {
	Text           string `xml:",chardata"`
	Idtf           string `xml:"idtf"`
	AssetNum       string `xml:"assetNum"`
	StopSeq        string `xml:"stopSeq"`
	Km             int    `xml:"km"`
	Pk             int    `xml:"pk"`
	M              int    `xml:"m"`
	TFTYPE         string `xml:"TF_TYPE"`
	TURN           string `xml:"TURN"`
	RADIUS         int    `xml:"RADIUS"`
	Number         int    `xml:"number"`
	Distance       int    `xml:"distance"`
	Zigzag         int    `xml:"zigzag"`
	Height         int    `xml:"height"`
	Offset         int    `xml:"offset"`
	Grounded       int    `xml:"grounded"`
	SPEED          int    `xml:"SPEED"`
	SuspensionType string `xml:"suspensionType"`
	Catenary       int    `xml:"catenary"`
	WireType       string `xml:"WireType"`
	CountWire      int    `xml:"CountWire"`
	Longitude      int    `xml:"longitude"`
	Latitude       int    `xml:"Latitude"`
	Gabarit        int    `xml:"Gabarit"`
}
