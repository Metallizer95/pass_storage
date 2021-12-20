package passport

type (
	UseCases interface {
		SavePassport() SavePassportUseCase
		LoadPassport() LoadPassportUseCase
	}

	SavePassportUseCase interface {
		Save(passport Model) *Model
	}

	LoadPassportUseCase interface {
		Load(id string) *Model
	}
)
