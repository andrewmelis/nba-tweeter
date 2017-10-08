package main

import (
	"net/http/httptest"
	"testing"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func TestNBA(t *testing.T) {
	now := clock.MakeTime(7, 30, "US/Eastern", "20170609")
	clock := clock.NewFakeClock(now)

	// setup fake server
	ts := httptest.NewServer(newFixtureHandlerFunc())
	defer ts.Close()
	url := newFakeScheduleURL(ts.URL) // fixme -- need path

	p := processor.NewDebugProcessor()
	w := nba.NewNBAWatcher(clock, p)
	s := nba.NewNBASchedule(url)

	t.Errorf("Found these games: %+v", s.Games())
	for _, g := range s.Games() {
		w.Follow(g)
	} // FIXME -- encapsulate this in a type?

	game := "GSWCLE"

	// game start
	expected := []string{"play 1", "play 2"}
	actual := p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance() // second quarter

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance() // third quarter

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance() // fourth quarter

	expected = []string{"play 1", "play 2"}
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	clock.Advance() // game over

	// expected is still same bc game over
	actual = p.Plays(game)
	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}
}
