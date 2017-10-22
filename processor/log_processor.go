package processor

import (
	"log"

	"github.com/andrewmelis/nba-tweeter/play"
)

type LogProcessor struct {
}

func NewLogProcessor() *LogProcessor {
	return &LogProcessor{}
}

func (p *LogProcessor) Process(gamecode string, play play.Play) {
	log.Printf("Processed play for %s:\n %s\n", gamecode, play)
}
