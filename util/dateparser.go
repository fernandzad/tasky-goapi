package util

import "time"

func StringDateParser(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
