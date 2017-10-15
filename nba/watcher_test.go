package nba

import (
	"testing"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/play"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func TestFollowProcessesEachPlayOfGameOnce(t *testing.T) {
	playsWithDup := []play.Play{
		fakePlay{"play 1"},
		fakePlay{"play 2"},
		fakePlay{"play 3"},
		fakePlay{"play 3"},
		fakePlay{"play 4"},
	}
	game := newFakeGame("GSWCLE", playsWithDup)

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p, func(string) {})

	w.Follow(game)
	c.Advance()

	expected := append(playsWithDup[:3], playsWithDup[4:]...) // delete duplicate
	actual := p.Plays(game.GameCode())

	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	for i, _ := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Wanted: %s, Got: %s", expected, actual)
		}
	}
}

func TestFollowProcessesPeriodically(t *testing.T) {
	plays := []play.Play{
		fakePlay{"play 1"},
		fakePlay{"play 2"},
		fakePlay{"play 3"},
		fakePlay{"play 4"},
	}

	game := newFakeGame("GSWCLE", plays)

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p, func(string) {})

	w.Follow(game)
	c.Advance()

	expected := plays
	actual := p.Plays(game.GameCode())

	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	for i, _ := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Wanted: %s, Got: %s", expected, actual)
		}
	}

	// add additional plays, simulating time passing
	newPlay := fakePlay{"play 5"}
	game.plays = append(game.plays, newPlay)

	c.Advance()

	expected = append(expected, newPlay)
	actual = p.Plays(game.GameCode())

	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	for i, _ := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Wanted: %s, Got: %s", expected, actual)
		}
	}
}

func TestStopsProcessesingWhenGameIsOver(t *testing.T) {
	plays := []play.Play{
		fakePlay{"play 1"},
		fakePlay{"play 2"},
		fakePlay{"play 3"},
		fakePlay{"play 4"},
	}

	game := newFakeGame("GSWCLE", plays)

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)
	cb := fakeCallback{}

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p, cb.cb)

	w.Follow(game)
	game.setActive(false)
	c.Advance()

	// it still gets all plays
	expected := plays
	actual := p.Plays(game.GameCode())

	if len(expected) != len(actual) {
		t.Errorf("Wanted: %s, Got: %s", expected, actual)
	}

	for i, _ := range actual {
		if expected[i] != actual[i] {
			t.Errorf("Wanted: %s, Got: %s", expected, actual)
		}
	}
	if cb.wasCleanupCalled != true {
		t.Errorf("Cleanup callback was not called")
	}
}

type fakeGame struct {
	code   string
	plays  []play.Play
	active bool
}

func newFakeGame(gamecode string, plays []play.Play) *fakeGame {
	return &fakeGame{gamecode, plays, true}
}

func (g *fakeGame) GameCode() string {
	return g.code
}

func (g *fakeGame) Plays() []play.Play {
	return g.plays
}

func (g *fakeGame) IsActive() bool {
	return g.active
}

func (g *fakeGame) setActive(b bool) {
	g.active = b
}

type fakePlay struct {
	s string
}

func (f fakePlay) String() string {
	return f.s
}

type fakeCallback struct {
	wasCleanupCalled bool
}

func (f *fakeCallback) cb(gamecode string) {
	f.wasCleanupCalled = true
}
