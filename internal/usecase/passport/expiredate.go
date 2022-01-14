package passport

import (
	"time"
)

// Time constants
const (
	timeLayout      = "02.01.2006 15:04:05" // Reference in time format.go
	hoursInDay      = 24
	daysInWeek      = 7
	halfExpireWeeks = 2
	expireWeeks     = 4
	warningTime     = time.Hour * hoursInDay * daysInWeek * halfExpireWeeks
	errorTime       = time.Hour * hoursInDay * daysInWeek * expireWeeks
)

// Status constants
const (
	StatusOk         = "OK"
	StatusSoonExpire = "Soon Expire"
	StatusExpired    = "Expired"
)

func IsExpired(s string) (bool, string) {
	return true, ""
}

func updateExpiration(m *Model) {
	changeDate := m.Header.CHANGEDATE
	t, _ := time.Parse(timeLayout, changeDate)
	now := time.Now()
	duration := now.Sub(t)
	var expirationModel ExpirationModel

	switch {
	case duration >= errorTime:
		expirationModel.Status = StatusExpired
		expirationModel.DaysUntilExpiration = "0"

	case duration >= warningTime:
		s := duration - warningTime
		expirationModel.DaysUntilExpiration = s.String()
		expirationModel.Status = StatusSoonExpire

	default:
		expirationModel.Status = StatusOk
		s := errorTime - duration
		expirationModel.DaysUntilExpiration = s.String()
	}

	m.Expiration = expirationModel
}
