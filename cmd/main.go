package main

import (
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/watcher"
)

func main() {
	s := nba.NewDefaultNBASchedule()

	// w := watcher.NewMultiWatcher(TweetWatcher, LogWatcher, DBWatcher) ???
	w := watcher.NewDebugWatcher()

	// loops over games
	// when new game
	// initialize game
	// game updates self with plays
	// watcher pipes all plays to a sink (e.g. twitter sink, log sink)

	w.Follow(s)
}
