package utility

import (
	"HELLO-GO/constant"
	"time"
)

func FormatDateString(date string, inputFormat string, outputFormat string) string {
	t, _ := time.Parse(inputFormat, date)
	return t.Format(outputFormat)
}

func GetTimeFromEpochAndZone(zone string, ts int64) time.Time {
	loc, err1 := time.LoadLocation(zone)

	var cTime time.Time
	secs := ts / 1000
	if err1 == nil {
		cTime = time.Unix(secs, 0).In(loc)
	} else {
		cTime = time.Unix(secs, 0)
	}

	return cTime
}

func GetTimeInFormattedString(t time.Time) string {
	// date format like
	return t.Format(constant.DT_FRMT_SEARCH_CNTXT)
}
