package passport

type Manager interface {
	SavePassport(passport Data) *Passport
	UpdatePassport(passport Passport) *Passport
	LoadPassportByID(id string) *Passport
	DeletePassportByID(id string) error
}

type passportManagerImpl struct {
	repo Repository
}

func (mng *passportManagerImpl) SavePassport(p Data) *Passport {
	return mng.repo.Create(p)
}

func (mng *passportManagerImpl) LoadPassportByID(id string) *Passport {
	return mng.repo.Read(id)
}

func (mng *passportManagerImpl) UpdatePassport(passport Passport) *Passport {
	return mng.repo.Update(passport)
}

func (mng *passportManagerImpl) DeletePassportByID(id string) error {
	return mng.repo.Delete(id)
}
