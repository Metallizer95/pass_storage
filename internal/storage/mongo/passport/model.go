package passportrepo

type RepositoryModel struct {
	Header HeaderRepoModel  `json:"header"`
	Towers []TowerRepoModel `json:"towers"`
}

type HeaderRepoModel struct {
	SiteID           string `json:"siteID"`
	SectionName      string `json:"sectionName"`
	SectionID        string `json:"sectionID"`
	EchName          string `json:"echName"`
	EchkName         string `json:"echkName"`
	Location         string `json:"location"`
	WayAmount        string `json:"wayAmount"`
	CurrentWay       string `json:"currentWay"`
	CurrentWayID     string `json:"currentWayID"`
	ChangeDate       string `json:"changeDate"`
	InitialMeter     string `json:"initialMeter"`
	InitialKm        string `json:"initialKm"`
	InitialPk        string `json:"initialPk"`
	InitialM         string `json:"initialM"`
	PlotLength       string `json:"plotLength"`
	SuspensionAmount string `json:"suspensionAmount"`
	WorkType         string `json:"workType"`
	Sequence         string `json:"sequence"`
}

type TowerRepoModel struct {
	ID             string `json:"id"`
	AssetNum       string `json:"assetNum"`
	StopSeq        string `json:"stopSeq"`
	Km             string `json:"km"`
	Pk             string `json:"pk"`
	M              string `json:"m"`
	Type           string `json:"type"`
	Turn           string `json:"turn"`
	Radius         string `json:"radius"`
	Number         string `json:"number"`
	Distance       string `json:"distance"`
	Zigzag         string `json:"zigzag"`
	Height         string `json:"height"`
	Offset         string `json:"offset"`
	Grounded       string `json:"grounded"`
	Speed          string `json:"speed"`
	SuspensionType string `json:"suspensionType"`
	Catenary       string `json:"catenary"`
	WireType       string `json:"wireType"`
	CountWire      string `json:"countWire"`
	Longitude      string `json:"longitude,omitempty"`
	Latitude       string `json:"latitude,omitempty"`
	Gabarit        string `json:"gabarit,omitempty"`
}

type ChangeDateCollectionModel struct {
	PassportID string `json:"passportID"`
	ChangeDate string `json:"ChangeDate"`
}
