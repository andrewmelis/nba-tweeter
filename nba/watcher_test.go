package nba

import (
	"testing"

	"github.com/andrewmelis/nba-tweeter/game"
)

func TestThisOnlyWorksForNBAGames(t *testing.T) {
	t.Error("is this the right design?")
}

func TestFollowStartsWatchingEveryGame(t *testing.T) {
	games := []string{"GSWCLE", "ATLWAS"}
	s := newFakeSchedule(games...)

	p := &Processor{}
	w := NewNBAWatcher(p)

	w.Follow(s)

	for _, game := range s.Games() {
		if !w.IsWatching(game) {
			t.Errorf("should be watching %s\n", game)
		}
	}
}

func TestFollowDoesSomething(t *testing.T) {
	t.Error("what does follow do? what side effect does it trigger?\n")
}

// func TestFollowProcessesEachPlayOfGame(t *testing.T) {
// 	games := []string{"GSWCLE"}
// 	s := newFakeSchedule(games...)

// 	p := &TestProcessor{}
// 	w := NewNBAWatcher(p)

// 	w.Follow(s)

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
