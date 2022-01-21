package passport

type Repository interface {
	Create(passport Data) *Passport
	CreateMany(passports []Data) error
	Read(id string) *Passport
	ReadAll() []Passport
	Update(passport Passport) *Passport
	Delete(id string) *Passport
}
