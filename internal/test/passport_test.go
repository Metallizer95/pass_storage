package test

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

var testFiles = []string{
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

	var inputPassportsModel []passport.Model

	for _, df := range testFiles {
		binData, err := ioutil.ReadFile(df)
		assert.NoError(t, err)

		var pm passport.Model
		assert.NoError(t, xml.Unmarshal(binData, &pm))
		inputPassportsModel = append(inputPassportsModel, pm)
		assert.NotNil(t, passportUseCases.SavePassportUseCase().Save(pm))
	}

	// Valid get passport
	for i := 0; i < len(inputPassportsModel); i++ {
		id := strconv.Itoa(i + 1)
		result := passportUseCases.LoadPassportUseCase().Load(id)
		assert.NotNil(t, result)

		comparePassportModels(t, inputPassportsModel[i], *result)
	}

	// Invalid get passport
	assert.Nil(t, passportUseCases.LoadPassportUseCase().Load("10000"))
}

func TestGetTowers(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}

func TestFindTowerByCoordinates(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}
