package main

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func TestNBA(t *testing.T) {
	setupNow()
	advanceTimeCh := setupTicker()

	ts := httptest.NewServer(newFixtureHandlerFunc())
	defer ts.Close()

	p := processor.NewDebugProcessor()
	w := nba.NewNBAWatcher(p, func(string) {})

	url := nba.NewNBAScheduleURL(ts.URL) //TODO
	s := nba.NewNBASchedule(url)

	f := nba.NewNBAFollower()

	f.Follow(s)

	for _, g := range s.Games() {
		w.Watch(g)
	} // FIXME -- encapsulate this in a type?

	game := "GSWCLE"
	if s.Games()[0].GameCode() != game {
		t.Errorf("Wanted: %s, Got: %s", game, s.Games()[0].GameCode())
	}

	// game start
	expected := []string{"play 1", "play 2"}
	actual := p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	advanceTimeCh <- nba.Now().Add(10 * time.Second) // second quarter

	time.Sleep(2 * time.Second)

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	advanceTimeCh <- nba.Now().Add(10 * time.Second) // third quarter
	time.Sleep(2 * time.Second)

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	advanceTimeCh <- nba.Now() // fourth quarter
	time.Sleep(2 * time.Second)

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	advanceTimeCh <- nba.Now() // game over
	time.Sleep(2 * time.Second)

	// expected is still same bc game over
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
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
	ticker := time.NewTicker(10 * time.Second)
	ch := make(chan time.Time, 10) // arbitrary buffer size
	ticker.C = ch

	nba.NewTicker = func(d time.Duration) *time.Ticker { return ticker }

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
