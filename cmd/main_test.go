package main

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/watcher"
)

func TestMain(t *testing.T) {
	// t.Errorf("blow up\n")
}

// FIXME for now only follow cavs games
func TestNBA(t *testing.T) {
	now := makeTime(7, 30, "US/Eastern", "20170402")
	clock := newFakeClock(now)

	// setup fake server
	ts := httptest.NewServer(newFixtureHandlerFunc("fixtures/example.json"))
	defer ts.Close()
	url := newFakeScheduleURL(ts.URL)

	p := processor.NewDebugProcessor()
	w := watcher.NewNbaWatcher(p)
	s := nba.NewNBASchedule(url)

	w.Follow(s)

	clock.Advance(30 * time.Second) // arbitrary number

	plays := []string{"play 1", "play 2"}

	expected := plays
	actual := w.Events()
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance(2 * time.Hour)

	expected = plays
	actual = w.Events()
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance(1 * time.Hour)

	// expected is still same bc game over
	actual = w.Events()
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}
}

// "feature test" -- run and actually tweet, then retrieve tweet via twitter api?
