package clock

import "time"

type Clock struct{}

func NewClock() *Clock {
	return &Clock{}
}

func (c *Clock) Now() time.Time {
	return time.Now()
}
