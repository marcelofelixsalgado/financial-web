package utils

import "time"

func ConvertStringToDateTime(source string) (time.Time, error) {
	layout := "02/01/2006T15:04:05Z"
	parsedDate, err := time.Parse(layout, source)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

// func ConvertDateTimeToString(source string) (time.Time, error) {
// 	layout := "2006-01-02T15:04:05Z"
// 	parsedDate, err := time.Parse(layout, source)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	return parsedDate, nil
// }
