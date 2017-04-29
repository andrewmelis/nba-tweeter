package watcher

import (
	"github.com/andrewmelis/nba-tweeter/game"
	"github.com/andrewmelis/nba-tweeter/schedule"
)

type Watcher interface {
	Follow(schedule.Schedule)
	IsWatching(game.Game) bool
}
