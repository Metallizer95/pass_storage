package common

type Manager interface {
	FindPassports(string) PassportsRoute
}

type manager struct {
	repository Repository
}

func NewManager(repository Repository) Manager {
	return &manager{repository: repository}
}

func (mng *manager) FindPassports(routeid string) PassportsRoute {
	return mng.repository.FindPassportsByRoute(routeid)
}
