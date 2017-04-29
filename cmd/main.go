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
	w.Follow(s)
}

type Watcher interface {
	Follow(schedule.Schedule)
	IsWatching(game.Game) bool
}

type PrintWatcher struct{}

func (w PrintWatcher) Follow(s schedule.Schedule) {
	fmt.Printf("%#v\n", s.Games())
}

func (w PrintWatcher) IsWatching(g game.Game) bool {
	return true
}
