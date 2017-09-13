package utc

import "time"

// Time in UTC *only*!
type Time struct {
	time time.Time
}

// GetTime as time.Time
func (t Time) GetTime() time.Time {
	return t.time
}

// NewTime to ensure UTC.
func NewTime(time time.Time) Time {
	return Time{time.UTC()}
}
