package errs

import (
	"encoding/xml"
	"errors"
)

var (
	ErrObjectNotFound      = errors.New("object not found in database")
	ErrObjectAlreadyExists = errors.New("object already exists")
	ErrNotFoundRoutes      = errors.New("not found routes in database")
)

type ErrorModel struct {
	XMLName xml.Name `xml:"Error"`
	Text    string   `xml:",chardata"`
	Message string   `xml:"message"`
}

func NewErrModel(err error) ErrorModel {
	return ErrorModel{Message: err.Error()}
}
