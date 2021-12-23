package passport

type (
	UseCases interface {
		SavePassportUseCase() SavePassportUseCase
		LoadPassportUseCase() LoadPassportUseCase
		GetTowersUseCase() GetTowersUseCase
	}

	SavePassportUseCase interface {
		Save(passport Model) *Model
	}

	LoadPassportUseCase interface {
		Load(id string) *Model
	}

	GetTowersUseCase interface {
		LoadTowers(id string) *TowersModel
	}
)

//TODO: По id паспорта вернуть все опоры участка - метод GET формат json
//TODO: По id участка и координате вернуть ближайшие опоры - метод GET, формат json
//TODO: По id участка и id опоры вернуть следующую и предыдущую опору - метод GET, формат json
