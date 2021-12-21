package routers

type Repository interface {
	Create(route ViksRoute) *ViksRoute
	Read(id string) *ViksRoute
	ReadAll() []ViksRoute
	Update(passport ViksRoute) *ViksRoute
	Delete(route ViksRoute) *ViksRoute
}
