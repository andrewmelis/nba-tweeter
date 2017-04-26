package schedule

import (
	"github.com/andrewmelis/nba-tweeter/game"
)

type Schedule interface {
	ScheduledGames() []game.Game
}
