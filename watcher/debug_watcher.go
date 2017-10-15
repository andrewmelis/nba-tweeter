package watcher

import (
	"fmt"

	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type DebugWatcher struct {
	plays []string
}

func NewDebugWatcher() *DebugWatcher {
	return &DebugWatcher{}
}

func (w *DebugWatcher) Watch(s schedule.Schedule) {
	fmt.Printf("%#v\n", s.Games())
	// TODO
}

func (w *DebugWatcher) IsWatching(g game.Game) bool {
	return true
}

func (w *DebugWatcher) Events() []string {
	return w.plays
}
