package passport

type Router struct {
	Passports []Passport
}

type Passport struct {
	ID string
	Data
}

type Data struct {
	Header
	Towers
}

type Header struct {
	SiteID           int
	SectionName      string
	SectionID        int
	EchName          string
	EchkName         string
	Location         string
	WayAmount        int
	CurrentWay       int
	CurrentWayID     int
	ChangeData       string
	InitialMeter     int
	InitialKm        int
	InitialPk        int
	InitialM         int
	PlotLength       int
	SuspensionAmount int
}

type Towers struct {
	Towers []Tower
}

type Tower struct {
	ID             string
	AssetNum       string
	StopSeq        string
	Km             int
	Pk             int
	M              int
	Type           string
	Turn           string
	Radius         int
	Number         int
	Distance       int
	Zigzag         int
	Height         int
	Offset         int
	Grounded       int
	Speed          int
	SuspensionType string
	Catenary       int
	WireType       string
	CountWire      int
	Longitude      int
	Latitude       int
	Gabarit        int
}
