package processor

import (
	"github.com/andrewmelis/nba-tweeter/play"
)

// Processor processes plays for games
type Processor interface {
	Process(string, play.Play)
}
