package models

import "time"

const (
	// TimeFormat is the format for Jira type "date-time".
	TimeFormat = "2006-01-02T15:04:05-0700"
	// DateFormat is the format for Jira type "date".
	DateFormat = "2006-01-02"
)

// DateScheme is a custom time type for Jira dates.
type DateScheme time.Time

// MarshalJSON marshals the DateScheme to JSON.
func (d *DateScheme) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format(DateFormat)), nil
}

// UnmarshalJSON unmarshals the DateScheme from JSON.
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

// MarshalJSON marshals the DateTimeScheme to JSON.
func (d *DateTimeScheme) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*d).Format(TimeFormat)), nil
}

// UnmarshalJSON unmarshals the DateTimeScheme from JSON.
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
