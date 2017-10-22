package main

import (
	"github.com/andrewmelis/nba-tweeter/nba"
	"github.com/andrewmelis/nba-tweeter/processor"
)

func main() {
	ch := make(chan struct{})

	s := nba.NewNBASchedule(nba.NewNBAScheduleURL("ljsdaf"))
	f := nba.NewNBAFollower(processor.ProcessorForGame)

	f.Follow(s)

	<-ch
}
