package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// t.Errorf("blow up\n")
}

type MockSchedule struct {
	Games []Game
}

type MockGame struct {
	Code string
}

func NewMockGame(code string) MockGame {
	return MockGame{Code: code}
}

// func (g MockGame) GameCode() string {
// 	// return g.GameCode
// 	return ""
// }

func NewMockSchedule(gameCodes []string) *MockSchedule {
	games := []Game{}
	for _, code := range gameCodes {
		games = append(games, NewMockGame(code))
	}
	return &MockSchedule{Games: games}
}

func (r *MockSchedule) ScheduledGames() []Game {
	return r.Games
}

type MockGameWatcher struct{}

func (w *MockGameWatcher) IsWatching(g Game) bool {
	return true
}

func (w *MockGameWatcher) Watch(g Game) error {
	return nil
}

func TestWatchOneGame(t *testing.T) {
	game := "CLEGSW"
	schedule := NewMockSchedule([]string{game})
	watcher := &MockGameWatcher{}

	watchGames(watcher, schedule)

	if len(schedule.ScheduledGames()) != 1 {
		t.Errorf("schedule not reflecting all games!")
	}

	for _, game := range schedule.ScheduledGames() {
		if !watcher.IsWatching(game) {
			t.Errorf("not watching %s!", game)
		}
	}
}
