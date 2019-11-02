// Package duration provides utilities for dealing with times and durations.
package duration

import (
	"time"

	"github.com/Cloud-Foundations/tricorder/go/tricorder/units"
)

// Duration represents a duration of time
// For negative durations, both Seconds and Nanoseconds are negative.
// Internal use only for now.
type Duration struct {
	Seconds     int64
	Nanoseconds int32
}

func New(d time.Duration) Duration {
	return newDuration(d)
}

// SinceEpoch returns the amount of time since unix epoch
func SinceEpoch(now time.Time) Duration {
	return sinceEpoch(now)
}

// SinceEpochFloat returns the amount of time since unix epoch
func SinceEpochFloat(secondsSinceEpoch float64) Duration {
	return sinceEpochFloat(secondsSinceEpoch)
}

// ParseWithUnit takes a string that is a quantity of unit and converts it
// to a Duration.
func ParseWithUnit(str string, unit units.Unit) (dur Duration, err error) {
	inUnit, err := parse(str)
	if err != nil {
		return
	}
	dur = inUnit.convert(unit, units.Second)
	return
}

// AsGoDuration converts this duration to a go duration
func (d Duration) AsGoDuration() time.Duration {
	return d.asGoDuration()
}

// AsGoTime Converts this duration to a go time in the
// system's local time zone.
func (d Duration) AsGoTime() time.Time {
	return d.asGoTime()
}

// AsFloat returns this duration in seconds.
func (d Duration) AsFloat() float64 {
	return d.asFloat()
}

// String shows in seconds
func (d Duration) String() string {
	return d.toString()
}

// StringUsingUnits shows in specified time unit.
// If unit not a time, shows in seconds.
func (d Duration) StringUsingUnits(unit units.Unit) string {
	return d.convert(units.Second, unit).toString()
}

// IsNegative returns true if this duration is negative.
func (d Duration) IsNegative() bool {
	return d.isNegative()
}

// PrettyFormat pretty formats this duration.
// PrettyFormat panics if this duration is negative.
func (d Duration) PrettyFormat() string {
	return d.prettyFormat()
}

// FloatToTime converts seconds after Jan 1, 1970 GMT to a time in the
// system's local time zone.
func FloatToTime(secondsSinceEpoch float64) time.Time {
	return SinceEpochFloat(secondsSinceEpoch).AsGoTime()
}

// TimeToFloat returns t as seconds after Jan 1, 1970 GMT
func TimeToFloat(t time.Time) (secondsSinceEpoch float64) {
	return SinceEpoch(t).AsFloat()
}

// ToFloat returns d as seconds
func ToFloat(d time.Duration) (seconds float64) {
	return float64(d) / float64(time.Second)
}

// FromFloat converts a value in seconds to a duration
func FromFloat(seconds float64) time.Duration {
	return time.Duration(seconds*float64(time.Second) + 0.5)
}
