package main

import (
	"fmt"
	"time"
)

func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	schedule, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return schedule
}

func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	dateFormatted, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	return dateFormatted.Before(time.Now())
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	appointment, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	hour := appointment.Hour()
	return hour >= 12 && hour <= 18
}

func Description(date string) string {
	layout := "Monday, January 2, 2006, at 15:04."
	schedule := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s", schedule.Format(layout))
}

func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
