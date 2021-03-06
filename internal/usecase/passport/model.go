package passport

import (
	"encoding/xml"
)

type Models struct {
	Model []Model
}

type Model struct {
	XMLName xml.Name    `xml:"SectionCertificate"`
	Text    string      `xml:",chardata"`
	ID      string      `xml:"id,omitempty" json:"id,omitempty"`
	Header  HeaderModel `xml:"Header" json:"Header"`
	Towers  TowersModel `xml:"Towers" json:"Towers"`
}

type OutputModel struct {
	Model
	Expiration ExpirationModel `xml:"Expiration" json:"Expiration"`
}

type HeaderModel struct {
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

type TowersModel struct {
	Text      string       `xml:",chardata"`
	SectionID string       `xml:"sectionID,omitempty" json:"sectionID,omitempty"`
	Towers    []TowerModel `xml:"Tower" json:"TowerModel"`
}

type TowerModel struct {
	Text           string `xml:",chardata"`
	Idtf           string `xml:"idtf"           json:"idtf"`
	AssetNum       string `xml:"assetNum"       json:"assetNum"`
	StopSeq        string `xml:"stopSeq"        json:"stopSeq"`
	Km             string `xml:"km"             json:"km"`
	Pk             string `xml:"pk"             json:"pk"`
	M              string `xml:"mapper"         json:"mapper"`
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
	Latitude       string `xml:"latitude"       json:"latitude"`
	Gabarit        string `xml:"Gabarit"        json:"Gabarit"`
}

type ExpirationModel struct {
	DaysUntilExpiration string `xml:"daysUntilExpiration" json:"daysUntilExpiration"`
	Status              string `xml:"status" json:"status"`
}

type ExpiredPassportsModel struct {
	XMLName   xml.Name       `xml:"ExpiredSectionSertificate"`
	Text      string         `xml:",chardata"`
	Passports []ExpiredModel `xml:"passports"`
}

type ExpiredModel struct {
	ID         string `xml:"id"`
	ChangeData string `xml:"CHANGEDATE"`
	Duration   string `xml:"duration"` // TODO: bad name of the field
}
