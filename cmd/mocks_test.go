package main

import (
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

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
