package main

import (
	"testing"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

func TestMain(t *testing.T) {
	// t.Errorf("blow up\n")
}

type MockSchedule struct {
	Games []game.Game
}

type MockGame struct {
	Code string
}

func NewMockGame(code string) MockGame {
	return MockGame{Code: code}
}

func (g MockGame) GameCode() string {
	// return g.GameCode
	return ""
}

func NewMockSchedule(gameCodes []string) *MockSchedule {
	games := []game.Game{}
	for _, code := range gameCodes {
		games = append(games, NewMockGame(code))
	}
	return &MockSchedule{Games: games}
}

func (r *MockSchedule) ScheduledGames() []game.Game {
	return r.Games
}

type MockGameWatcher struct{}

func (w *MockGameWatcher) IsWatching(g game.Game) bool {
	return true
}

func (w *MockGameWatcher) Follow(s schedule.Schedule) error {
	return nil
}

func TestWatchOneGame(t *testing.T) {
	game := "CLEGSW"
	schedule := NewMockSchedule([]string{game})
	watcher := &MockGameWatcher{}

	watcher.Follow(schedule)

	if len(schedule.ScheduledGames()) != 1 {
		t.Errorf("schedule not reflecting all games!")
	}

	for _, game := range schedule.ScheduledGames() {
		if !watcher.IsWatching(game) {
			t.Errorf("not watching %s!", game)
		}
	}
}

// "feature test" -- run and actually tweet, then retrieve tweet via twitter api?
