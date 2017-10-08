package clock

import (
	"time"
)

func MakeTime(pmHour, minute int, location, date string) time.Time {
	l, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, time.April, 2, 12+pmHour, 0, 0, 0, l) // known cavs game
	return t
}

type FakeClock struct {
	now time.Time
}

func NewFakeClock(startTime time.Time) *FakeClock {
	return &FakeClock{startTime}
}

func (c *FakeClock) Now() time.Time {
	return c.now
}

func (c *FakeClock) Advance() {
	d := 10 * time.Minute // arbitrary
	c.now = c.now.Add(d)
}
