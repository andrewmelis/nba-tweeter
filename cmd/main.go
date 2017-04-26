package main

import (
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

func main() {
	// s := NewNBASchedule()
	// w := NewNBAWatcher()
}

type GameWatcher interface {
	Follow(schedule.Schedule)
	IsWatching(game.Game) bool
}

func watch(w GameWatcher, s schedule.Schedule) {
	w.Follow(s)
}
