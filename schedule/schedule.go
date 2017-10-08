package schedule

import (
	"github.com/andrewmelis/nba-tweeter/game"
)

// Schedule returns all active games
type Schedule interface {
	Games() []game.Game
}
