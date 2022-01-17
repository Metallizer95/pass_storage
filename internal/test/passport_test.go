package test

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"store_server/pkg/logging"

	passport2 "store_server/internal/domain/passport"
	"store_server/internal/usecase/passport"
	"testing"
)

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
	_, teardown := TestDatabase(t)
	defer teardown()
}

func TestGetTowers(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}

func TestFindTowerByCoordinates(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}
