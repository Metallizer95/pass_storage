package test

import "testing"

var testRouteFile = "route_template/route_template.xml"

func TestSaveRouteSeveralTimes(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}

func TestGetRoute(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}

func TestGetPassportsRoute(t *testing.T) {
	_, teardown := TestDatabase(t)
	defer teardown()
}
