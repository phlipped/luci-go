// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package settings

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

// A collection of useful templating functions

// funcMap is what gets fed into the template bundle.
var funcMap = template.FuncMap{
	"humanDuration": humanDuration,
	"humanTimeRFC":  humanTimeRFC,
	"startswith":    strings.HasPrefix,
	"sub":           sub,
	"shortHash":     shortHash,
}

// humanDuration takes a time t in seconds as a duration and translates it
// into a human readable string of x units y units, where x and y could be in
// days, hours, minutes, or seconds, whichever is the largest.
func humanDuration(t uint64) string {
	// Input: Duration in seconds.  Output, the duration pretty printed.
	day := t / 86400
	hr := (t % 86400) / 3600
	min := (t % 3600) / 60
	sec := t % 60

	if day > 0 {
		if hr != 0 {
			return fmt.Sprintf("%d days %d hrs", day, hr)
		}
		return fmt.Sprintf("%d days", day)
	} else if hr > 0 {
		if min != 0 {
			return fmt.Sprintf("%d hrs %d mins", hr, min)
		}
		return fmt.Sprintf("%d hrs", hr)
	} else {
		if min > 0 {
			if sec != 0 {
				return fmt.Sprintf("%d mins %d secs", min, sec)
			}
			return fmt.Sprintf("%d mins", min)
		}
		return fmt.Sprintf("%d secs", sec)
	}
}

// humanTimeRFC takes in the time represented as a RFC3339 string and returns
// something more human readable (like RFC850: Monday, 02-Jan-06 15:04:05 MST).
func humanTimeRFC(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t, err = time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return s
		}
	}
	return t.Format(time.RFC850)
}

// sub subtracts one number from another, because apperently go templates aren't
// smart enough to do that.
func sub(a, b int) int {
	return a - b
}

// shortHash abbriviates a git hash into 6 characters.
func shortHash(s string) string {
	if len(s) > 6 {
		return s[0:6]
	}
	return s
}
