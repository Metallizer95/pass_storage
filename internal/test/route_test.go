package test

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	passport2 "store_server/internal/domain/passport"
	routers2 "store_server/internal/domain/routers"
	"store_server/internal/usecase/passport"
	"store_server/internal/usecase/routers"
	"testing"
)

var (
	testRouteFile   = "route_template/route_template.xml"
	amountPassports = 3
)

func TestSaveRouteSeveralTimes(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	_ = saveTestPassports(t, passportUseCases.SavePassportUseCase(), testPassportFiles)

	routeManager := routers2.NewRouteManager(db.RouteRepository())
	routeUseCases := routers.NewUseCases(routeManager, passportUseCases)

	// Save route
	dataFile, err := ioutil.ReadFile(testRouteFile)
	assert.NoError(t, err)

	var inputRouteModel routers.RouteModel
	assert.NoError(t, xml.Unmarshal(dataFile, &inputRouteModel))

	assert.NotNil(t, routeUseCases.SaveRouter().Save(inputRouteModel))
	assert.Nil(t, routeUseCases.SaveRouter().Save(inputRouteModel))
}

func TestGetRoute(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	_ = saveTestPassports(t, passportUseCases.SavePassportUseCase(), testPassportFiles)

	routeManager := routers2.NewRouteManager(db.RouteRepository())
	routeUseCases := routers.NewUseCases(routeManager, passportUseCases)

	dataFile, err := ioutil.ReadFile(testRouteFile)
	assert.NoError(t, err)

	var inputRouteModel routers.RouteModel
	assert.NoError(t, xml.Unmarshal(dataFile, &inputRouteModel))
	assert.NotNil(t, routeUseCases.SaveRouter().Save(inputRouteModel))

	gotRoute := routeUseCases.LoadRouterByID().Load(inputRouteModel.ViksRouteID)
	assert.NotNil(t, gotRoute)

	compareRoutes(t, inputRouteModel, *gotRoute)

	gotAllRoutes := routeUseCases.LoadRouters().Load()
	assert.NotNil(t, gotAllRoutes)
	assert.Equal(t, 1, len(gotAllRoutes.RouteModel))
}

func TestGetPassportsRoute(t *testing.T) {
	db, teardown := TestDatabase(t)
	defer teardown()

	passportManager := passport2.NewPassportManager(db.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	passportsModel := saveTestPassports(t, passportUseCases.SavePassportUseCase(), testPassportFiles)

	routeManager := routers2.NewRouteManager(db.RouteRepository())
	routeUseCases := routers.NewUseCases(routeManager, passportUseCases)

	dataFile, err := ioutil.ReadFile(testRouteFile)
	assert.NoError(t, err)

	var inputRouteModel routers.RouteModel
	assert.NoError(t, xml.Unmarshal(dataFile, &inputRouteModel))

	assert.NotNil(t, routeUseCases.SaveRouter().Save(inputRouteModel))

	ps := routeUseCases.LoadPassportsByRoute().Load(inputRouteModel.ViksRouteID)
	assert.NotNil(t, ps)
	assert.Equal(t, len(ps.Passports), amountPassports)

	for i := 0; i < amountPassports; i++ {
		comparePassportModels(t, passportsModel[0], ps.Passports[0])
	}

}
