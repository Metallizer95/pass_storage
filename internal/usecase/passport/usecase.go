package passport

type SavePassportUseCase interface {
	Save(passport PassportModel) error
}

type SavePassportUseCaseImpl struct {
}

func (s *SavePassportUseCaseImpl) Save(passport PassportModel) error {
	return nil
}
