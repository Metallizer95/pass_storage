package routers

type Manager interface {
	SaveRoute(route ViksRoute) *ViksRoute
	LoadRouteByID(id string) *ViksRoute
	LoadRoutes() []ViksRoute
	UpdateRoute(route ViksRoute) *ViksRoute
	DeleteRoute(route ViksRoute) *ViksRoute
}

type manager struct {
	repo Repository
}

func NewRouteManager(repo Repository) Manager {
	return &manager{repo: repo}
}

func (mng *manager) SaveRoute(route ViksRoute) *ViksRoute {
	return mng.repo.Create(route)
}

func (mng *manager) LoadRouteByID(id string) *ViksRoute {
	return mng.repo.Read(id)
}

func (mng *manager) LoadRoutes() []ViksRoute {
	return mng.repo.ReadAll()
}

func (mng *manager) UpdateRoute(route ViksRoute) *ViksRoute {
	return mng.repo.Update(route)
}

func (mng *manager) DeleteRoute(route ViksRoute) *ViksRoute {
	return mng.repo.Delete(route)
}
