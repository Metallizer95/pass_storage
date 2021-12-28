package passport

type Passport struct {
	ID string
	Data
}

type Data struct {
	Header
	Towers
}

type Header struct {
	SiteID           string
	SectionName      string
	SectionID        string
	EchName          string
	EchkName         string
	Location         string
	WayAmount        string
	CurrentWay       string
	CurrentWayID     string
	ChangeDate       string
	InitialMeter     string
	InitialKm        string
	InitialPk        string
	InitialM         string
	PlotLength       string
	SuspensionAmount string
	WorkType         string
	Sequence         string
}

type Towers struct {
	Towers map[string]Tower
}

type Tower struct {
	ID             string
	AssetNum       string
	StopSeq        string
	Km             string
	Pk             string
	M              string
	Type           string
	Turn           string
	Radius         string
	Number         string
	Distance       string
	Zigzag         string
	Height         string
	Offset         string
	Grounded       string
	Speed          string
	SuspensionType string
	Catenary       string
	WireType       string
	CountWire      string
	Longitude      string
	Latitude       string
	Gabarit        string
}
