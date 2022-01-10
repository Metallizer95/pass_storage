package routerepo

type repositoryModel struct {
	ID    string               `bson:"_id"   json:"id"`
	Route routeRepositoryModel `bson:"route" json:"route"`
}

type routeRepositoryModel struct {
	MasterPMNum    string   `json:"MasterPMNum"    bson:"MasterPMNum"`
	TripChangeDate string   `json:"TripChangeDate" bson:"TripChangeDate"`
	TripType       string   `json:"TripType"       bson:"TripType"`
	ViksRoutedID   string   `json:"ViksRoutedID"   bson:"ViksRoutedID"`
	Car            string   `json:"Car"            bson:"Car"`
	CarID          string   `json:"CarID"          bson:"CarID"`
	Description    string   `json:"Description"    bson:"Description"`
	Eigthnum       string   `json:"Eigthnum"       bson:"Eigthnum"`
	SectionSet     []string `json:"SectionSet"     bson:"SectionSet"`
}
