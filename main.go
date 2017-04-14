package main

import (
	_ "fmt"
)

func main() {
	// for all games
	// watch games
	// tweet each play
}

type Game interface {
	// GameCode() string
}

type Schedule interface {
	ScheduledGames() []Game
}

type GameWatcher interface {
	Watch(Game) error
	IsWatching(Game) bool
}

func watchGames(w GameWatcher, s Schedule) error {
	for _, game := range s.ScheduledGames() {
		w.Watch(game) // TODO do something with errors
	}
	return nil
}
