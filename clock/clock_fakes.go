package clock

import (
	"time"
)

func MakeTime(timeString, location string) time.Time {
	l, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}

	var layout = "20060102 3:04pm"
	t, err := time.ParseInLocation(layout, timeString, l)
	if err != nil {
		panic(err)
	}
	return t
}

type FakeClock struct {
	now time.Time
	c   chan time.Time
}

func NewFakeClock(startTime time.Time) *FakeClock {
	return &FakeClock{startTime, make(chan time.Time)}
}

func (c *FakeClock) Now() time.Time {
	return c.now
}

func (c *FakeClock) Ticker() <-chan time.Time {
	return c.c
}

func (c *FakeClock) Advance() {
	d := 10 * time.Minute // arbitrary
	c.now = c.now.Add(d)
	select {
	case c.c <- c.Now():
		time.Sleep(100 * time.Millisecond) // let goroutines do stuff in testing
	case <-time.After(100 * time.Millisecond):
		// timeout
	}
}
