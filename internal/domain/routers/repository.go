package routers

type Repository interface {
	Create(passport Data) *Passport
	Read(id string) *Passport
	Update(passport Passport) *Passport
	Delete(id string) *Passport
}
