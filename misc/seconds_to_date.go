package misc

import (
	"fmt"
	"strings"
)

func SecondsToDate(seconds int) string {
	const (
		minute = 60
		hour   = 3600
		day    = 86400
		year   = 31536000
	)

	var sol string

	if seconds >= year {
		years := seconds / year
		sol = fmt.Sprintf("%s %v years,", sol, years)

		seconds = seconds % year
	}

	if len(sol) > 0 || seconds >= day {
		days := seconds / day
		sol = fmt.Sprintf("%s %v days,", sol, days)

		seconds = seconds % day
	}

	if len(sol) > 0 || seconds >= hour {
		hours := seconds / hour
		sol = fmt.Sprintf("%s %v hours,", sol, hours)

		seconds = seconds % hour
	}

	if len(sol) > 0 || seconds >= minute {
		minutes := seconds / minute
		sol = fmt.Sprintf("%s %v minutes,", sol, minutes)

		seconds = seconds % minute
	}

	return strings.TrimSpace(fmt.Sprintf("%s %v seconds", sol, seconds))
}
