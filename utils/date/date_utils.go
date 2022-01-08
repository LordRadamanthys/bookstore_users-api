package date

import "time"

const (
	apiDateLayout = "2006-01-T15:04:05Z"
)

func GetDateNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
