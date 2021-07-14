package schema

import "testing"

func TestFormatChecker(t *testing.T) {
	data := map[string]string{
		"23:12:07":                        "time",
		"2006-01-02":                      "date",
		"02 Jan 06 15:04 MST":             "date-time",
		"02 Jan 06 15:04 -0700":           "date-time",
		"Mon, 02 Jan 2006 15:04:05 MST":   "date-time",
		"Mon, 02 Jan 2006 15:04:05 -0700": "date-time",
		"Monday, 02-Jan-06 15:04:05 MST":  "date-time",
		"w@dasd.asd":                      "email",
	}
	for s, expected := range data {
		t.Run(s, func(t *testing.T) {
			format := formatChecker(s)
			if format != expected {
				t.Errorf("Format of \"%v\" was incorrect, got: \"%v\", want: %v.", s, format, expected)
			}
		})

	}
}
