package test

import (
	"context"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"store_server/internal/storage/mongorepo"
	"store_server/internal/usecase/passport"
	"testing"
)

func TestDatabase(t *testing.T) (mongorepo.Client, func()) {
	client, err := newClient()
	if err != nil {
		t.Fatal(err)
	}
	return client, func() {
		_ = client.db.Database(dbName).Drop(context.TODO())
	}
}

func saveTestPassports(t *testing.T, useCase passport.SavePassportUseCase, paths []string) []passport.Model {
	t.Helper()

	var result []passport.Model
	for _, path := range paths {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error(err)
		}

		var inputPassportsModel passport.Model
		err = xml.Unmarshal(file, &inputPassportsModel)
		if err != nil {
			t.Error(err)
		}
		r := *useCase.Save(inputPassportsModel)
		assert.NotNil(t, r)

		result = append(result, r)
	}
	return result
}

func comparePassportModels(t *testing.T, p1, p2 passport.Model) {
	t.Helper()
	assert.Equal(t, p1.Header.SectionName, p2.Header.SectionName)
	assert.Equal(t, p1.Header.CHANGEDATE, p2.Header.CHANGEDATE)
	assert.Equal(t, p1.Header.SectionID, p2.Header.SectionID)
	assert.Equal(t, p1.Header.SiteID, p2.Header.SiteID)
	assert.Equal(t, p1.Header.Sequence, p2.Header.Sequence)
	assert.Equal(t, p1.Header.WorkType, p2.Header.WorkType)
	assert.Equal(t, p1.Header.SuspensionAmount, p2.Header.SuspensionAmount)
	assert.Equal(t, p1.Header.PlotLength, p2.Header.PlotLength)
	assert.Equal(t, p1.Header.InitialM, p2.Header.InitialM)
	assert.Equal(t, p1.Header.EchkName, p2.Header.EchkName)
	assert.Equal(t, p1.Header.InitialPK, p2.Header.InitialPK)
	assert.Equal(t, p1.Header.LocationId, p2.Header.LocationId)
	assert.Equal(t, p1.Header.InitialKm, p2.Header.InitialKm)
	assert.Equal(t, p1.Header.InitialMeter, p2.Header.InitialMeter)
	assert.Equal(t, p1.Header.CurrentWayID, p2.Header.CurrentWayID)
	assert.Equal(t, p1.Header.WayAmount, p2.Header.WayAmount)
	assert.Equal(t, p1.Header.EchName, p2.Header.EchName)
	compareTowers(t, p1.Towers, p2.Towers)
}

func compareTowers(t *testing.T, t1, t2 passport.TowersModel) {
	t.Helper()
	for n, tower := range t1.Towers {
		tower2 := t2.Towers[n]

		assert.Equal(t, tower.Idtf, tower2.Idtf)
		assert.Equal(t, tower.M, tower2.M)
		assert.Equal(t, tower.Gabarit, tower2.Gabarit)
		assert.Equal(t, tower.Latitude, tower2.Latitude)
		assert.Equal(t, tower.Longitude, tower2.Longitude)
		assert.Equal(t, tower.CountWire, tower2.CountWire)
		assert.Equal(t, tower.WireType, tower2.WireType)
		assert.Equal(t, tower.Catenary, tower2.Catenary)
		assert.Equal(t, tower.SuspensionType, tower2.SuspensionType)
		assert.Equal(t, tower.AssetNum, tower2.AssetNum)
		assert.Equal(t, tower.Grounded, tower2.Grounded)
		assert.Equal(t, tower.Offset, tower2.Offset)
		assert.Equal(t, tower.Height, tower2.Height)
		assert.Equal(t, tower.Zigzag, tower2.Zigzag)
		assert.Equal(t, tower.Distance, tower2.Distance)
		assert.Equal(t, tower.Number, tower2.Number)
		assert.Equal(t, tower.Pk, tower2.Pk)
		assert.Equal(t, tower.Km, tower2.Km)
		assert.Equal(t, tower.StopSeq, tower2.StopSeq)
		assert.Equal(t, tower.RADIUS, tower2.RADIUS)
		assert.Equal(t, tower.SPEED, tower2.SPEED)
		assert.Equal(t, tower.TFTYPE, tower2.TFTYPE)
		assert.Equal(t, tower.TURN, tower2.TURN)
	}
}
