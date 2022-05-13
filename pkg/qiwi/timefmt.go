package qiwi

import (
	"fmt"
	"time"
)

type QiwiTime string

// ConvetToQiwiTime - converts golang time object to specific format required by the qiwi api
func ConvertToQiwiTime(t time.Time) QiwiTime {
	var qiwiTime QiwiTime

	// formatting time
	qiwiTimeWithoutTimeZone := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day()+1, t.Hour(), t.Minute(), t.Second())

	// adding time zone
	_, offsetInSeconds := t.Zone()
	offsetInHours := offsetInSeconds / 3600
	if offsetInHours > 0 {
		qiwiTime = QiwiTime(qiwiTimeWithoutTimeZone + fmt.Sprintf("+%02d:00", offsetInHours))
	} else {
		qiwiTime = QiwiTime(qiwiTimeWithoutTimeZone + fmt.Sprintf("-%02d:00", offsetInHours))
	}

	return qiwiTime
}
