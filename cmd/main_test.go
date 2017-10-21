package main

import (
	_ "fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func TestNBA(t *testing.T) {
	setupNow()
	advanceTimeCh := setupTicker()

	ts := httptest.NewServer(newFixtureHandlerFunc())
	defer ts.Close()

	nba.BaseURL = ts.URL
	url := nba.NewNBAScheduleURL(ts.URL) //TODO
	s := nba.NewNBASchedule(url)

	spy := NewSpyProcessorFactory()
	f := nba.NewNBAFollower(spy.MakeProcessor)

	hookCh := make(chan struct{})
	nba.WatchHook = func() {
		<-hookCh
	}

	f.Follow(s)

	game := "GSWCLE"
	if s.Games()[0].GameCode() != game {
		t.Errorf("Wanted: %s, Got: %s", game, s.Games()[0].GameCode())
	}

	advanceTimeCh <- nba.Now().Add(10 * time.Second)
	hookCh <- struct{}{}

	p := spy.ps[game]

	expectedQ1LenPlays := 156
	expected := expectedQ1LenPlays
	actual := p.Plays(game)
	if expected != len(actual) {
		t.Errorf("Wanted: %d, Got: %d", expected, len(actual))
	}

	hookCh <- struct{}{}

	expectedQ2LenPlays := 114
	expected += expectedQ2LenPlays
	actual = p.Plays(game)
	if expected != len(actual) {
		t.Errorf("Wanted: %d, Got: %d", expected, len(actual))
	}

	hookCh <- struct{}{}

	expectedQ3LenPlays := 112
	expected += expectedQ3LenPlays
	actual = p.Plays(game)
	if expected != len(actual) {
		t.Errorf("Wanted: %d, Got: %d", expected, len(actual))
	}

	hookCh <- struct{}{}

	expectedQ4LenPlays := 122
	expected += expectedQ4LenPlays
	actual = p.Plays(game)
	if expected != len(actual) {
		t.Errorf("Wanted: %d, Got: %d", expected, len(actual))
	}

	hookCh <- struct{}{}

	// expected is still same bc game over
	actual = p.Plays(game)
	if expected != len(actual) {
		t.Errorf("Wanted: %d, Got: %d", expected, len(actual))
	}
}

func setupNow() {
	now := makeTime("20170609 7:30pm", "US/Eastern")
	nba.Now = func() time.Time { return now }
}

// setupTicker allow test to control when time advances with ch <- time.Now()
// if multiple callers have asked for ticker,
// need to send down channel multiple times -- probably some way to multiplex
func setupTicker() chan time.Time {
	ticker := time.NewTicker(100 * time.Millisecond)
	ch := make(chan time.Time, 10) // arbitrary buffer size
	ticker.C = ch

	var wasDebugTickerTaken bool
	nba.NewTicker = func(d time.Duration) *time.Ticker {
		if !wasDebugTickerTaken {
			wasDebugTickerTaken = true
			// fmt.Printf("returning debug ticker %#v\n", ticker)
			return ticker
		}
		newTicker := time.NewTicker(100 * time.Millisecond)
		// fmt.Printf("returning new ticker %#v\n", newTicker)
		return newTicker
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

type SpyProcessorFactory struct {
	ps map[string]*processor.DebugProcessor
}

func NewSpyProcessorFactory() *SpyProcessorFactory {
	return &SpyProcessorFactory{
		make(map[string]*processor.DebugProcessor),
	}
}

func (f *SpyProcessorFactory) MakeProcessor(g game.Game) processor.Processor {
	p := processor.NewDebugProcessor()
	f.ps[g.GameCode()] = p
	return p
}
