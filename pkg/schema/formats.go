package schema

import (
	"net/mail"
	"time"
)

type stringFormat struct {
	name  string
	match func(string) bool
}

func formatChecker(s string) string {
	formats := []stringFormat{
		{"email", emailCheck},
		{"date-time", dateTimeCheck},
		{"date", dateCheck},
		{"time", timeCheck},
	}
	for _, format := range formats {
		if format.match(s) {
			return format.name
		}
	}
	return s
}

func emailCheck(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func dateTimeCheck(s string) bool {
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		time.RFC822,
		time.RFC822Z,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC850,
	}

	for _, format := range formats {
		if _, err := time.Parse(format, s); err == nil {
			return true
		}
	}

	return false
}

func dateCheck(s string) bool {
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}

func timeCheck(s string) bool {
	_, err := time.Parse("15:04:05", s)
	return err == nil
}
