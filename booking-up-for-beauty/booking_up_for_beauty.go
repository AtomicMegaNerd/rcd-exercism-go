package booking

import (
	"fmt"
	"time"
)

const shortForm = "1/2/2006 15:04:05"
const longForm = "January 2, 2006 15:04:05"
const longerYetForm = "Monday, January 2, 2006 15:04:05"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse(shortForm, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	now := time.Now()
	t, _ := time.Parse(longForm, date)
	return t.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(longerYetForm, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t, _ := time.Parse(shortForm, date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(2022, 9, 15, 0, 0, 0, 0, time.UTC)
}
