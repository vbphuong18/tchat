package helper

import "time"

const FormatDateTime = "2006-01-02 15:04:05"

var (
	LocLocal *time.Location
)

func ConvertStringToTime(timeStr string) (time.Time, error) {
	t, err := time.ParseInLocation(FormatDateTime, timeStr, LocLocal)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func init() {
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	LocLocal = l
}
