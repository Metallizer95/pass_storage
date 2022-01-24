package tests

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"store_server/pkg/logging"
	"strconv"

	passport2 "store_server/internal/domain/passport"
	"store_server/internal/usecase/passport"
	"testing"
)

var testPassportFiles = []string{
	"passport_templates/pt_seq1.xml",
	"passport_templates/pt_seq2.xml",
	"passport_templates/pt_seq3.xml",
}

func TestMain(m *testing.M) {
	logging.New(false, os.Stdout)
	m.Run()
}

func TestSavePassportSeveralTimes(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	dataFile, err := ioutil.ReadFile("passport_templates/pt_seq1.xml")
	assert.NoError(t, err)

	// Save passport
	var inputPassportModel passport.Model
	assert.NoError(t, xml.Unmarshal(dataFile, &inputPassportModel))
	assert.NotNil(t, passportUseCases.SavePassportUseCase().Save(inputPassportModel))

	// Save the same passport and expect nil result
	assert.Nil(t, passportUseCases.SavePassportUseCase().Save(inputPassportModel))
}

func TestGetPassportsByID(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	inputPassportsModel := saveTestPassports(t, passportUseCases.SavePassportUseCase(), testPassportFiles)

	// Valid get passport
	for i := 0; i < len(inputPassportsModel); i++ {
		id := strconv.Itoa(i + 1)
		result := passportUseCases.LoadPassportUseCase().Load(id)
		assert.NotNil(t, result)

		comparePassportModels(t, inputPassportsModel[i].Model, result.Model)
	}

	// Invalid get passport
	assert.Nil(t, passportUseCases.LoadPassportUseCase().Load("10000"))
}

func TestGetTowers(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	dataFile, err := ioutil.ReadFile("passport_templates/pt_seq1.xml")
	assert.NoError(t, err)

	// Save passport
	var inputPassportModel passport.Model
	assert.NoError(t, xml.Unmarshal(dataFile, &inputPassportModel))
	assert.NotNil(t, passportUseCases.SavePassportUseCase().Save(inputPassportModel))

	towers := passportUseCases.GetTowersUseCase().LoadAllTowerByPassportId(inputPassportModel.Header.SectionID)
	assert.NotNil(t, towers)
	assert.Equal(t, len(inputPassportModel.Towers.Towers), len(towers.Towers))
	compareTowers(t, inputPassportModel.Towers, *towers)
}

func TestFindTowerByCoordinates(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	dataFile, err := ioutil.ReadFile("passport_templates/pt_seq1.xml")
	assert.NoError(t, err)

	// Save passport
	var inputPassportModel passport.Model
	assert.NoError(t, xml.Unmarshal(dataFile, &inputPassportModel))
	assert.NotNil(t, passportUseCases.SavePassportUseCase().Save(inputPassportModel))

	testCases := []struct {
		longitude         float64
		latitude          float64
		expectedLongitude string
		expectedLatitude  string
	}{
		{
			longitude:         9.0,
			latitude:          8.0,
			expectedLongitude: "6.5",
			expectedLatitude:  "8",
		},
		{
			longitude:         16.0,
			latitude:          15.0,
			expectedLongitude: "12",
			expectedLatitude:  "14",
		},
		{
			longitude:         1.5,
			latitude:          2.0,
			expectedLongitude: "1",
			expectedLatitude:  "2",
		},
		{
			longitude:         7.0,
			latitude:          7.3,
			expectedLongitude: "6.5",
			expectedLatitude:  "8",
		},
	}

	for _, tc := range testCases {
		tow := passportUseCases.FindTowerByIdAndCoordinateUseCase().FindTower(inputPassportModel.Header.SectionID, tc.longitude, tc.latitude)
		assert.NotNil(t, tow)
		assert.Equal(t, tc.expectedLatitude, tow.Latitude)
		assert.Equal(t, tc.expectedLongitude, tow.Longitude)
	}
}
