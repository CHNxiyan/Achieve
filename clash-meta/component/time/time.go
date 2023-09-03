package time

import (
	"net/http"
	"strconv"
	"time"
)

var offset time.Duration
var enable = true

func init() {
	if b, _ := strconv.ParseBool("DISABLE_TIME_AUTO_SYNC"); b {
		enable = false
	}
}

func SaveTimeFromHttpHeader(header http.Header) {
	now := time.Now()
	if t, err := time.Parse(http.TimeFormat, header.Get("Date")); err == nil {
		if d := t.Sub(now); d.Abs() > 30*time.Second {
			offset = d
		} else {
			offset = 0
		}
	}
}

func Now() time.Time {
	now := time.Now()
	if enable && offset.Abs() > 0 {
		now = now.Add(offset)
	}
	return now
}
