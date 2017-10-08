package watcher

import (
	"github.com/andrewmelis/nba-tweeter/game"
)

// Watcher follows games
// Watcher ensures every play is processed in a watched game
// Watcher is responsible for only processing each play once
type Watcher interface {
	Follow(game.Game) // should return err?
}

// TODO!!!!!!!!!
// <Something> ensures that each active game in the followed schedule
// is being watched. It checks for active games every <period>.
