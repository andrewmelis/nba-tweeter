package nba

import (
	"fmt"
	"time"
)

type ScheduleURL interface {
	URL() string
}

type NBAScheduleURL struct{}

func NewDefaultNBAScheduleURL() ScheduleURL {
	return NBAScheduleURL{}
}

func (u NBAScheduleURL) URL() string {
	// TODO retrieve from "today url":
	// https://data.nba.net/10s/prod/v1/today.json -> currentScoreboard
	now := time.Now()
	date := now.Format("20060102")
	fmt.Println(date)
	return fmt.Sprintf("https://data.nba.net/data/10s/prod/v1/%s/scoreboard.json", date)
}
