package main

import (
	"path/filepath"
	"regexp"
	"time"
)

var dateMarkPattern = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

// StrSliceContains slice of string contains
func StrSliceContains(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}

// BeginningOfDay beginning of day
func BeginningOfDay() time.Time {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.UTC)
}

// FindDateMark find date mark from filename
func FindDateMark(file string) (t time.Time, ok bool) {
	// check filename, search for date-mark
	var match []string
	if match = dateMarkPattern.FindStringSubmatch(filepath.Base(file)); len(match) != 1 {
		return
	}
	// decode date-mark, skip if invalid
	var err error
	if t, err = time.Parse("2006-01-02", match[0]); err != nil {
		return
	}
	// k
	ok = true
	return
}
