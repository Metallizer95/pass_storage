package routers

type Repository interface {
	Create(route Router) *Router
	Read(id string) *Router
	Update(passport Router) *Router
	Delete(id string) *Router
}
