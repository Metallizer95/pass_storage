package passportctrl

import "errors"

var (
	ErrNotFindLongitude            = errors.New("not find parameter longitude")
	ErrNotFindLatitude             = errors.New("not find parameter latitude")
	ErrWrongTypeLongitudeParameter = errors.New("wrong type parameter longitude")
	ErrWrongTypeLatitudeParameter  = errors.New("wrong type parameter latitude")
)
