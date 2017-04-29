package main

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

func main() {
	s := nba.NewDefaultNBASchedule()
	w := PrintWatcher{}
	watch(w, s)

	// fmt.Printf("%s\n", s.ScheduledGames())
	// w := NewNBAWatcher()
}

type Watcher interface {
	Follow(schedule.Schedule)
	IsWatching(game.Game) bool
}

func watch(w Watcher, s schedule.Schedule) {
	w.Follow(s)
}

type PrintWatcher struct {
	s schedule.Schedule
}

func (w PrintWatcher) Follow(s schedule.Schedule) {
	fmt.Printf("%#v\n", s.Games())
}

func (w PrintWatcher) IsWatching(g game.Game) bool {
	return true
}
