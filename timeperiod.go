package techan

import (
	"regexp"
	"time"
)

// TimePeriod is a simple struct that describes a period of time with a Start and End time
type TimePeriod struct {
	Start time.Time
	End   time.Time
}

// Constants representing basic, human-readable and writable date formats
const (
	SimpleDateTimeFormat = "01/02/2006T15:04:05"
	SimpleDateFormat     = "01/02/2006"

	SimpleTimeFormat   = "15:04:05"
	SimpleDateFormatV2 = "2006-01-02"
)

// Constants representing regexes for parsing datetimes
var (
	SimpleTimeFomatRegex    = regexp.MustCompile(`T\d{2}:\d{2}:\d{2}`)
	SimpleDateFormatV2Regex = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
)

// ParseTimePeriod parses two datetimes as one string and returns it as a TimePeriod.
//
// Note that if you were previously using Parse, the date format has changed to something more rfc3339-like (yyyy-mm-dd)
// Will accept any combination of date and time for either side. Omitting the right hand side will result in a time
// period ending in time.Now()
func ParseTimePeriod(period string) (TimePeriod, error) {
	_ = "STUB: not implemented"
	return *new(TimePeriod), nil
}

// Parse takes a string in one of the following formats and returns a new TimePeriod, and optionally, an error
//
// Deprecated: Please use ParseTimePeriod instead
func Parse(timerange string) (tr TimePeriod, err error) {
	_ = "STUB: not implemented"
	return *new(TimePeriod), nil
}

// In returns a copy of TimePeriod tp with both start and end times' location set to the specified location
func (tp TimePeriod) In(location *time.Location) TimePeriod {
	_ = "STUB: not implemented"
	return *new(TimePeriod)
}

// UTC returns a copy of TimePeriod tp with both start and end times' location set to UTC
func (tp TimePeriod) UTC() TimePeriod {
	_ = "STUB: not implemented"
	return *

	// Length returns the length of the period as a time.Duration value
	new(TimePeriod)
}

func (tp TimePeriod) Length() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Since returns the amount of time elapsed since the end of another TimePeriod as a time.Duration value
func (tp TimePeriod) Since(other TimePeriod) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Format returns the string representation of this timePeriod in the given format
func (tp TimePeriod) Format(layout string) string { _ = "STUB: not implemented"; return "" }

// Advance will return a new TimePeriod with the start and end periods moved forwards or backwards in time in accordance
// with the number of iterations given.
//
// Example:
// A timePeriod that is one hour long, starting at unix time 0 and ending at unix time 3600, and advanced by one,
// will return a time period starting at unix time 3600 and ending at unix time 7200
func (tp TimePeriod) Advance(iterations int) TimePeriod {
	_ = "STUB: not implemented"
	return *new(TimePeriod)
}

func (tp TimePeriod) String() string { _ = "STUB: not implemented"; return "" }

// NewTimePeriod returns a TimePeriod starting at the given time and ending at the given time plus the given duration
func NewTimePeriod(start time.Time, period time.Duration) TimePeriod {
	_ = "STUB: not implemented"
	return *new(TimePeriod)
}
