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
	game := &fakeGame{"GSWCLE", playsWithDup}

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p)

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

func TestFollowProcessesMultipleGames(t *testing.T) {
	playsG1 := []play.Play{
		fakePlay{"play 1"},
		fakePlay{"play 2"},
	}

	playsG2 := make([]play.Play, len(playsG1))
	copy(playsG2, playsG1)

	game1 := &fakeGame{"GSWCLE", playsG1}
	game2 := &fakeGame{"OKCSAS", playsG2}

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p)

	w.Follow(game1)
	w.Follow(game2)
	c.Advance()

	games := []*fakeGame{game1, game2}
	plays := [][]play.Play{playsG1, playsG2}
	for i, _ := range games {
		expected := plays[i]
		game := games[i].GameCode()
		actual := p.Plays(games[i].GameCode())

		if len(expected) != len(actual) {
			t.Errorf("Game: %s, Wanted: %s, Got: %s", game, expected, actual)
		}

		for i, _ := range actual {
			if expected[i] != actual[i] {
				t.Errorf("Game: %s, Wanted: %s, Got: %s", game, expected, actual)
			}
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

	game := &fakeGame{"GSWCLE", plays}

	now := clock.MakeTime("20170609 7:30pm", "US/Eastern")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p)

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
	t.Errorf("how to clean up?")
}

type fakeGame struct {
	code  string
	plays []play.Play
}

func (g *fakeGame) GameCode() string {
	return g.code
}

func (g *fakeGame) Plays() []play.Play {
	return g.plays
}

type fakePlay struct {
	s string
}

func (f fakePlay) String() string {
	return f.s
}
