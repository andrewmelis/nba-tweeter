package main

import (
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/watcher"
)

func main() {
	s := nba.NewDefaultNBASchedule()
	w := watcher.NewDebugWatcher()
	w.Follow(s)
}
