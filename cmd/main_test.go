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
	MockGames []game.Game
}

type MockGame struct {
	Code string
}

func NewMockGame(code string) MockGame {
	return MockGame{Code: code}
}

func (g MockGame) GameCode() string {
	return ""
}

func NewMockSchedule(gameCodes []string) *MockSchedule {
	games := []game.Game{}
	for _, code := range gameCodes {
		games = append(games, NewMockGame(code))
	}
	return &MockSchedule{games}
}

func (r *MockSchedule) Games() []game.Game {
	return r.MockGames
}

type MockWatcher struct{}

func (w *MockWatcher) IsWatching(g game.Game) bool {
	return true
}

func (w *MockWatcher) Follow(s schedule.Schedule) error {
	return nil
}

func TestWatchOneGame(t *testing.T) {
	game := "CLEGSW"
	schedule := NewMockSchedule([]string{game})
	watcher := &MockWatcher{}

	watcher.Follow(schedule)

	if len(schedule.Games()) != 1 {
		t.Errorf("schedule not reflecting all games!")
	}

	for _, game := range schedule.Games() {
		if !watcher.IsWatching(game) {
			t.Errorf("not watching %s!", game)
		}
	}
}

// "feature test" -- run and actually tweet, then retrieve tweet via twitter api?
