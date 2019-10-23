package file

import (
	"strings"
	"time"
)

var (
	m = map[string]string{
		"%Y": "2006",
		"%M": "01",
		"%D": "02",
		"%h": "15",
		"%m": "04",
		"%s": "05",
	}
)

// WithTime formats string by time.
func WithTime(str string, t ...time.Time) string {
	var now time.Time
	if len(t) == 0 {
		now = time.Now()
	} else {
		now = t[1]
	}
	for k, v := range m {
		str = strings.Replace(str, k, now.Format(v), -1)
	}
	return str
}
