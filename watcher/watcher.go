package watcher

import (
	"github.com/andrewmelis/nba-tweeter/game"
)

// Watcher follows a game
// Watcher ensures every play is processed in a watched game
// Watcher is responsible for only processing each play once
type Watcher interface {
	Watch(game.Game) // should return err?
}
