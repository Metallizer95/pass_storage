package passport

type (
	UseCases interface {
		SavePassportUseCase() SavePassportUseCase
		LoadPassportUseCase() LoadPassportUseCase
		GetTowersUseCase() GetTowersUseCase
		FindTowerByIdAndCoordinateUseCase() FindTowerByIdAndCoordinateUseCase
	}

	SavePassportUseCase interface {
		Save(passport Model) *Model
	}

	LoadPassportUseCase interface {
		Load(id string) *Model
	}

	GetTowersUseCase interface {
		LoadAllTowerByPassportId(passportId string) *TowersModel
		LoadTowerById(passportId, towerId string) *TowerModel
	}

	FindTowerByIdAndCoordinateUseCase interface {
		FindTower(id string, longitude float64, latitude float64) *TowerModel
	}
)

//TODO: По id участка и координате вернуть ближайшие опоры - метод GET, формат json
//TODO: По id участка и id опоры вернуть следующую и предыдущую опору - метод GET, формат json
