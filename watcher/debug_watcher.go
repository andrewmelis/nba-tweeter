package watcher

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type DebugWatcher struct{}

func NewDebugWatcher() DebugWatcher {
	return DebugWatcher{}
}

func (w DebugWatcher) Follow(s schedule.Schedule) {
	fmt.Printf("%#v\n", s.Games())
}

func (w DebugWatcher) IsWatching(g game.Game) bool {
	return true
}
