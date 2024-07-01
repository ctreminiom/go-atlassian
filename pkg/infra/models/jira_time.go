package models

import "time"

const (
	TimeFormat = "2006-01-02T15:04:05-0700"
	DateFormat = "2006-01-02"
)

// DateScheme is a custom time type for Jira dates.
type DateScheme time.Time

func (d *DateScheme) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format(DateFormat)), nil
}

func (d *DateScheme) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	parsed, err := time.Parse(`"`+DateFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*d = DateScheme(parsed)
	return nil
}

// DateTimeScheme is a custom time type for Jira times.
type DateTimeScheme time.Time

func (d *DateTimeScheme) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format(TimeFormat)), nil
}

func (d *DateTimeScheme) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	parsed, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	if err != nil {
		return err
	}

	*d = DateTimeScheme(parsed)
	return nil
}
