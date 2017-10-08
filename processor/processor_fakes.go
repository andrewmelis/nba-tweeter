package processor

import (
	"github.com/andrewmelis/nba-tweeter/play"
)

type DebugProcessor struct {
	plays map[string][]play.Play
}

func NewDebugProcessor() *DebugProcessor {
	plays := make(map[string][]play.Play, 0)
	return &DebugProcessor{plays}
}

func (p *DebugProcessor) Process(gamecode string, play play.Play) {
	p.plays[gamecode] = append(p.plays[gamecode], play)
}

func (p *DebugProcessor) Plays(gamecode string) []play.Play {
	return p.plays[gamecode]
}
