package passport

import (
	"encoding/xml"
)

type Model struct {
	XMLName xml.Name `xml:"SectionCertificate"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,omitempty" json:"id,omitempty"`
	Header  Header   `xml:"Header" json:"Header"`
	Towers  Towers   `xml:"Towers" json:"Towers"`
}

type Header struct {
	Text             string `xml:",chardata"`
	SiteID           string `xml:"siteId"           json:"siteId"`
	SectionName      string `xml:"sectionName"      json:"sectionName"`
	SectionID        string `xml:"sectionId"        json:"sectionId"`
	EchName          string `xml:"echName"          json:"echName"`
	EchkName         string `xml:"echkName"         json:"echkName"`
	LocationId       string `xml:"locationId"       json:"locationId"`
	WayAmount        string `xml:"wayAmount"        json:"wayAmount"`
	CurrentWay       string `xml:"currentWay"       json:"currentWay"`
	CurrentWayID     string `xml:"currentWayID"     json:"currentWayID"`
	CHANGEDATE       string `xml:"CHANGEDATE"       json:"CHANGEDATE"`
	InitialMeter     string `xml:"initialMeter"     json:"initialMeter"`
	InitialKm        string `xml:"initialKM"        json:"initialKM"`
	InitialPK        string `xml:"initialPK"        json:"initialPK"`
	InitialM         string `xml:"initialM"         json:"initialM"`
	PlotLength       string `xml:"plotLength"       json:"plotLength"`
	SuspensionAmount string `xml:"suspensionAmount" json:"suspensionAmount"`
	Sequence         string `xml:"Sequence"         json:"Sequence"`
	WorkType         string `xml:"workType"         json:"workType"`
}

type Towers struct {
	Text   string  `xml:",chardata"`
	Towers []Tower `xml:"Tower" json:"Tower"`
}

type Tower struct {
	Text           string `xml:",chardata"`
	Idtf           string `xml:"idtf"           json:"idtf"`
	AssetNum       string `xml:"assetNum"       json:"assetNum"`
	StopSeq        string `xml:"stopSeq"        json:"stopSeq"`
	Km             string `xml:"km"             json:"km"`
	Pk             string `xml:"pk"             json:"pk"`
	M              string `xml:"m"              json:"m"`
	TFTYPE         string `xml:"TF_TYPE"        json:"TF_TYPE"`
	TURN           string `xml:"TURN"           json:"TURN"`
	RADIUS         string `xml:"RADIUS"         json:"RADIUS"`
	Number         string `xml:"number"         json:"number"`
	Distance       string `xml:"distance"       json:"distance"`
	Zigzag         string `xml:"zigzag"         json:"zigzag"`
	Height         string `xml:"height"         json:"height"`
	Offset         string `xml:"offset"         json:"offset"`
	Grounded       string `xml:"grounded"       json:"grounded"`
	SPEED          string `xml:"SPEED"          json:"SPEED"`
	SuspensionType string `xml:"suspensionType" json:"suspensionType"`
	Catenary       string `xml:"catenary"       json:"catenary"`
	WireType       string `xml:"WireType"       json:"WireType"`
	CountWire      string `xml:"CountWire"      json:"CountWire"`
	Longitude      string `xml:"longitude"      json:"longitude"`
	Latitude       string `xml:"Latitude"       json:"Latitude"`
	Gabarit        string `xml:"Gabarit"        json:"Gabarit"`
}
