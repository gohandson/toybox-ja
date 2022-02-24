package eventcal

import "time"

var DefaultClock Clock = ClockFunc(time.Now)

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}
