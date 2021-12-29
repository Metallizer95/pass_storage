package passport

import "time"

const (
	timeLayout  = "02.01.2006 15:04:05" // Reference in time format.go
	hoursInDay  = 24
	daysInWeek  = 7
	expireWeeks = 3
	deadline    = time.Hour * hoursInDay * daysInWeek * expireWeeks
)

func IsExpired(changeDate string) (bool, string) {
	t, _ := time.Parse(timeLayout, changeDate)
	now := time.Now()
	duration := now.Sub(t)
	if duration > deadline {
		return true, duration.String()
	}
	return false, ""
}
