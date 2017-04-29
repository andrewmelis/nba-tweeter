package main

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

func main() {
	s := nba.NewDefaultNBASchedule()
	w := PrintGameWatcher{}
	watch(w, s)

	// fmt.Printf("%s\n", s.ScheduledGames())
	// w := NewNBAWatcher()
}

type GameWatcher interface {
	Follow(schedule.Schedule)
	IsWatching(game.Game) bool
}

func watch(w GameWatcher, s schedule.Schedule) {
	w.Follow(s)
}

type PrintGameWatcher struct {
	s schedule.Schedule
}

func (w PrintGameWatcher) Follow(s schedule.Schedule) {
	fmt.Printf("%#v\n", s.ScheduledGames())
}

func (w PrintGameWatcher) IsWatching(g game.Game) bool {
	return true
}
