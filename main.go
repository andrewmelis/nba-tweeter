package main

import (
	_ "fmt"
)

func main() {
	// loop here?

	// for all games
	// watch games
	// tweet each play
}

type Game interface {
	// GameCode() string
}

type Schedule interface {
	ScheduledGames() []Game // should this return a channel instead?
}

type GameWatcher interface {
	Follow(Schedule)
	IsWatching(Game) bool
}
