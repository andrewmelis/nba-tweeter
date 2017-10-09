package game

import (
	"github.com/andrewmelis/nba-tweeter/play"
)

// Game represents a match between two teams
// A game retrieves its own plays
type Game interface {
	GameCode() string // TODO make gamecode a type?
	Plays() []play.Play
	IsActive() bool
}
