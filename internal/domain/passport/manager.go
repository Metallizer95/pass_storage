package passport

type Manager interface {
	SavePassport(passport Data) *Passport
	UpdatePassport(passport Passport) *Passport
	LoadPassportByID(id string) *Passport
	LoadAll() []Passport
	DeletePassportByID(id string) *Passport
}

type passportManagerImpl struct {
	repo Repository
}

func NewPassportManager(repository Repository) Manager {
	return &passportManagerImpl{repo: repository}
}

func (mng *passportManagerImpl) SavePassport(p Data) *Passport {
	return mng.repo.Create(p)
}

func (mng *passportManagerImpl) LoadPassportByID(id string) *Passport {
	return mng.repo.Read(id)
}

func (mng *passportManagerImpl) LoadAll() []Passport {
	return mng.repo.ReadAll()
}

func (mng *passportManagerImpl) UpdatePassport(passport Passport) *Passport {
	return mng.repo.Update(passport)
}

func (mng *passportManagerImpl) DeletePassportByID(id string) *Passport {
	return mng.repo.Delete(id)
}
