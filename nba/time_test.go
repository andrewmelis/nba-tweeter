package nba

import (
	"time"
)

func setupNow() {
	now := makeTime("20170609 7:30pm", "US/Eastern")
	Now = func() time.Time { return now }
}

// setupTicker allow test to control when time advances with ch <- time.Now()
// if multiple callers have asked for ticker,
// need to send down channel multiple times -- probably some way to multiplex
func setupTicker() chan time.Time {
	ticker := time.NewTicker(10 * time.Second)
	ch := make(chan time.Time, 10) // arbitrary buffer size
	ticker.C = ch

	var wasDebugTickerTaken bool

	NewTicker = func(d time.Duration) *time.Ticker {
		if !wasDebugTickerTaken {
			wasDebugTickerTaken = true
			return ticker
		}
		return time.NewTicker(10 * time.Second)
	}

	return ch
}

func makeTime(timeString, location string) time.Time {
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
