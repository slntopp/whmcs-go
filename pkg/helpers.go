package whmcs

import "time"

type FormattedTime string

func Time(t time.Time) FormattedTime {
	return FormattedTime(t.Format(time.DateTime))
}

// Returns parsed time as golang type
func (t FormattedTime) Parse() time.Time {
	time, _ := time.Parse(time.DateTime, string(t))
	return time
}

type Boolean float64

func Bool(b bool) Boolean {
	if b {
		return 1
	}
	return 0
}

func (b Boolean) Bool() bool {
	return b != 0
}
