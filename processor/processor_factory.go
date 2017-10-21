package processor

import (
	"github.com/andrewmelis/nba-tweeter/game"
)

type ProcessorFactory func(game.Game) Processor

func ProcessorForGame(g game.Game) Processor {
	return NewDebugProcessor()
}
