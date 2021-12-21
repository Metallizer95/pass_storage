package passport

import (
	"encoding/xml"
)

type Model struct {
	XMLName xml.Name `xml:"SectionCertificate"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,omitempty"`
	Header  Header   `xml:"Header"`
	Towers  Towers   `xml:"Towers"`
}

type Header struct {
	Text             string `xml:",chardata"`
	SiteID           string `xml:"siteId"`
	SectionName      string `xml:"sectionName"`
	SectionID        string `xml:"sectionId"`
	EchName          string `xml:"echName"`
	EchkName         string `xml:"echkName"`
	LocationId       string `xml:"locationId"`
	WayAmount        string `xml:"wayAmount"`
	CurrentWay       string `xml:"currentWay"`
	CurrentWayID     string `xml:"currentWayID"`
	CHANGEDATA       string `xml:"CHANGEDATA"`
	InitialMeter     string `xml:"initialMeter"`
	InitialKm        string `xml:"initialKM"`
	InitialPK        string `xml:"initialPK"`
	InitialM         string `xml:"initialM"`
	PlotLength       string `xml:"plotLength"`
	SuspensionAmount string `xml:"suspensionAmount"`
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
	Km             string `xml:"km"`
	Pk             string `xml:"pk"`
	M              string `xml:"m"`
	TFTYPE         string `xml:"TF_TYPE"`
	TURN           string `xml:"TURN"`
	RADIUS         string `xml:"RADIUS"`
	Number         string `xml:"number"`
	Distance       string `xml:"distance"`
	Zigzag         string `xml:"zigzag"`
	Height         string `xml:"height"`
	Offset         string `xml:"offset"`
	Grounded       string `xml:"grounded"`
	SPEED          string `xml:"SPEED"`
	SuspensionType string `xml:"suspensionType"`
	Catenary       string `xml:"catenary"`
	WireType       string `xml:"WireType"`
	CountWire      string `xml:"CountWire"`
	Longitude      string `xml:"longitude"`
	Latitude       string `xml:"Latitude"`
	Gabarit        string `xml:"Gabarit"`
}
