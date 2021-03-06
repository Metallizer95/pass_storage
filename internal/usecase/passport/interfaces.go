package passport

type (
	UseCases interface {
		SavePassportUseCase() SavePassportUseCase
		LoadPassportUseCase() LoadPassportUseCase
		GetTowersUseCase() GetTowersUseCase
		FindTowerByIdAndCoordinateUseCase() FindTowerByIdAndCoordinateUseCase
		FindExpiredPassportsUseCase() FindExpiredPassportsUseCase
	}

	SavePassportUseCase interface {
		Save(passport Model) *OutputModel
		SaveMany(passports []Model) error
	}

	LoadPassportUseCase interface {
		Load(id string) *OutputModel
	}

	GetTowersUseCase interface {
		LoadAllTowerByPassportId(passportId string) *TowersModel
		LoadTowerById(passportId, towerId string) *TowerModel
	}

	FindTowerByIdAndCoordinateUseCase interface {
		FindTower(id string, longitude float64, latitude float64) *TowerModel
	}

	FindExpiredPassportsUseCase interface {
		FindPassports() ExpiredPassportsModel
	}
)
