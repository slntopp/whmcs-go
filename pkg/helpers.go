package whmcs

import "time"

type formattedTime string

func Time(t time.Time) *formattedTime {
	time := formattedTime(t.Format(time.DateTime))
	return &time
}

// Returns parsed time as golang type
func (t *formattedTime) Parse() time.Time {
	time, _ := time.Parse(time.DateTime, string(*t))
	return time
}

type boolean int

func BoolInt(b bool) *boolean {
	res := boolean(0)
	if b {
		res = 1
	}
	return &res
}

func (b *boolean) Bool() bool {
	return *b != 0
}

func Bool(v bool) *bool {
	return &v
}

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}
