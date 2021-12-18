package passport

type Router struct {
	Passports []Model `xml:"passports"`
}

type Model struct {
	Header `xml:"Header"`
	Towers `xml:"Towers"`
}

//TODO: Which format does change data have?
//TODO: What is M?

type Header struct {
	SiteID           int    `xml:"siteID"`
	SectionName      string `xml:"sectionName"`
	SectionID        int    `xml:"sectionID"`
	EchName          string `xml:"echName"`
	EchkName         string `xml:"echkName"`
	Location         string `xml:"locationId"`
	WayAmount        int    `xml:"wayAmount"`
	CurrentWay       int    `xml:"currentWay"`
	CurrentWayID     int64  `xml:"currentWayID"`
	ChangeData       string `xml:"CHANGEDATA"`
	InitialMeter     int64  `xml:"initialMeter"`
	InitialKm        int    `xml:"initialKM"`
	InitialPk        int    `xml:"initialPK"`
	InitialM         int    `xml:"initialM"`
	PlotLength       int    `xml:"plotLength"`
	SuspensionAmount int    `xml:"suspensionAmount"`
}

type Towers struct {
	Towers []Tower `xml:"Tower"`
}

type Tower struct {
	ID             string `xml:"idtf"`
	AssetNum       string `xml:"assetNum"`
	StopSeq        string `xml:"stopSeq"`
	Km             int    `xml:"km"`
	Pk             int    `xml:"pk"`
	M              int    `xml:"m"`
	Type           string `xml:"TF_TYPE,omitempty"`
	Turn           string `xml:"TURN"`
	Radius         int    `xml:"RADIUS"`
	Number         int    `xml:"number"`
	Distance       int    `xml:"distance"`
	Zigzag         int    `xml:"zigzag"`
	Height         int    `xml:"height"`
	Offset         int    `xml:"offset"`
	Grounded       int    `xml:"grounded"`
	Speed          int    `xml:"SPEED"`
	SuspensionType string `xml:"suspensionType"`
	Catenary       int    `xml:"catenary"`
	WireType       string `xml:"WireType"`
	CountWire      int    `xml:"CountWire"`
	Longitude      int    `xml:"longitude,omitempty"`
	Latitude       int    `xml:"Latitude,omitempty"`
	Gabarit        int    `xml:"Gabarit,omitempty"`
}
