package test

import "testing"

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
