package passportrepo

type RepositoryModel struct {
	ID     string           `bson:"id"`
	Header HeaderRepoModel  `bson:"header"`
	Towers []TowerRepoModel `bson:"towers"`
}

type HeaderRepoModel struct {
	SiteID           string `bson:"siteID"`
	SectionName      string `bson:"sectionName"`
	SectionID        string `bson:"sectionID"`
	EchName          string `bson:"echName"`
	EchkName         string `bson:"echkName"`
	Location         string `bson:"location"`
	WayAmount        string `bson:"wayAmount"`
	CurrentWay       string `bson:"currentWay"`
	CurrentWayID     string `bson:"currentWayID"`
	ChangeDate       string `bson:"changeDate"`
	InitialMeter     string `bson:"initialMeter"`
	InitialKm        string `bson:"initialKm"`
	InitialPk        string `bson:"initialPk"`
	InitialM         string `bson:"initialM"`
	PlotLength       string `bson:"plotLength"`
	SuspensionAmount string `bson:"suspensionAmount"`
	WorkType         string `bson:"workType"`
	Sequence         string `bson:"sequence"`
}

type TowerRepoModel struct {
	ID             string `bson:"id"`
	AssetNum       string `bson:"assetNum"`
	StopSeq        string `bson:"stopSeq"`
	Km             string `bson:"km"`
	Pk             string `bson:"pk"`
	M              string `bson:"m"`
	Type           string `bson:"type"`
	Turn           string `bson:"turn"`
	Radius         string `bson:"radius"`
	Number         string `bson:"number"`
	Distance       string `bson:"distance"`
	Zigzag         string `bson:"zigzag"`
	Height         string `bson:"height"`
	Offset         string `bson:"offset"`
	Grounded       string `bson:"grounded"`
	Speed          string `bson:"speed"`
	SuspensionType string `bson:"suspensionType"`
	Catenary       string `bson:"catenary"`
	WireType       string `bson:"wireType"`
	CountWire      string `bson:"countWire"`
	Longitude      string `bson:"longitude,omitempty"`
	Latitude       string `bson:"latitude,omitempty"`
	Gabarit        string `bson:"gabarit,omitempty"`
}

type ChangeDateCollectionModel struct {
	PassportID string `bson:"passportID"`
	ChangeDate string `bson:"ChangeDate"`
}
