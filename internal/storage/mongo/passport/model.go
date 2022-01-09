package passportrepo

type RepositoryModel struct {
	ID     string           `bson:"_id"    json:"id"`
	Header HeaderRepoModel  `bson:"header" json:"header"`
	Towers []TowerRepoModel `bson:"towers" json:"towers"`
}

type HeaderRepoModel struct {
	SiteID           string `bson:"siteID"             json:"siteID"`
	SectionName      string `bson:"sectionName"        json:"sectionName"`
	SectionID        string `bson:"sectionID"          json:"sectionID"`
	EchName          string `bson:"echName"            json:"echName"`
	EchkName         string `bson:"echkName"           json:"echkName"`
	Location         string `bson:"location"           json:"location"`
	WayAmount        string `bson:"wayAmount"          json:"wayAmount"`
	CurrentWay       string `bson:"currentWay"         json:"currentWay"`
	CurrentWayID     string `bson:"currentWayID"       json:"currentWayID"`
	ChangeDate       string `bson:"changeDate"         json:"changeDate"`
	InitialMeter     string `bson:"initialMeter"       json:"initialMeter"`
	InitialKm        string `bson:"initialKm"          json:"initialKm"`
	InitialPk        string `bson:"initialPk"          json:"initialPk"`
	InitialM         string `bson:"initialM"           json:"initialM"`
	PlotLength       string `bson:"plotLength"         json:"plotLength"`
	SuspensionAmount string `bson:"suspensionAmount"   json:"suspensionAmount"`
	WorkType         string `bson:"workType"           json:"workType"`
	Sequence         string `bson:"sequence"           json:"sequence"`
}

type TowerRepoModel struct {
	ID             string `bson:"id"                     json:"id"`
	AssetNum       string `bson:"assetNum"               json:"assetNum"`
	StopSeq        string `bson:"stopSeq"                json:"stopSeq"`
	Km             string `bson:"km"                     json:"km"`
	Pk             string `bson:"pk"                     json:"pk"`
	M              string `bson:"m"                      json:"m"`
	Type           string `bson:"type"                   json:"type"`
	Turn           string `bson:"turn"                   json:"turn"`
	Radius         string `bson:"radius"                 json:"radius"`
	Number         string `bson:"number"                 json:"number"`
	Distance       string `bson:"distance"               json:"distance"`
	Zigzag         string `bson:"zigzag"                 json:"zigzag"`
	Height         string `bson:"height"                 json:"height"`
	Offset         string `bson:"offset"                 json:"offset"`
	Grounded       string `bson:"grounded"               json:"grounded"`
	Speed          string `bson:"speed"                  json:"speed"`
	SuspensionType string `bson:"suspensionType"         json:"suspensionType"`
	Catenary       string `bson:"catenary"               json:"catenary"`
	WireType       string `bson:"wireType"               json:"wireType"`
	CountWire      string `bson:"countWire"              json:"countWire"`
	Longitude      string `bson:"longitude,omitempty"    json:"longitude,omitempty"`
	Latitude       string `bson:"latitude,omitempty"     json:"latitude,omitempty"`
	Gabarit        string `bson:"gabarit,omitempty"      json:"gabarit,omitempty"`
}

type ChangeDateCollectionModel struct {
	PassportID string `bson:"_id"        json:"id"`
	ChangeDate string `bson:"ChangeDate" json:"ChangeDate"`
}
