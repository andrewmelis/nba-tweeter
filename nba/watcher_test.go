package nba

import (
	"testing"

	"github.com/andrewmelis/nba-tweeter/clock"
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func TestThisOnlyWorksForNBAGames(t *testing.T) {
	t.Error("is this the right design?")
}

func TestFollowProcessesEachPlayOfGame(t *testing.T) {
	games := []string{"GSWCLE", "ATLWAS"}
	s := newFakeSchedule(games...)

	now := clock.MakeTime(7, 30, "US/Eastern", "20170609")
	c := clock.NewFakeClock(now)

	p := processor.NewDebugProcessor()
	w := NewNBAWatcher(c, p)

	w.Follow(s)

	for _, game := range s.Games() {
		if !w.IsWatching(game) {
			t.Errorf("should be watching %s\n", game)
		}
	}
}

// 	// advance "time" by one play
// 	s.Games[0].AdvanceTime()

// 	// check if first play of game was processed
// 	firstPlays := s.Games[0].Plays[0]
// 	if !p.Games(s.Games[0]).WasProcessed(s.Games[0].Plays[0]) {
// 		t.Errorf("first play %s of game %s should have been processed",
// 	}

// 	// advance "time" by another play
// 	// check if second play of game was processed
// }

type fakeSchedule struct {
	games []fakeGame
}

func newFakeSchedule(gameCodes ...string) fakeSchedule {
	var s fakeSchedule
	for _, g := range gameCodes {
		s.games = append(s.games, fakeGame{g})
	}
	return s
}

func (s fakeSchedule) Games() []game.Game {
	var games []game.Game
	for _, g := range s.games {
		games = append(games, g)
	}
	return games
}

type fakeGame struct {
	code string
}

func (g fakeGame) GameCode() string {
	return g.code
}
