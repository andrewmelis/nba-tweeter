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

	/*
		now := time.Now() FIXME
		date := now.Format("20060102")
	*/

	// vvvvvvvvvvvvvvvvvvvvvvvv
	raw, err := time.Parse("20060102", "20170717")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	date := raw.Format("20060102")
	// ^^^^^^^^^^^^^^^^^^^^^^^^

	fmt.Println(date)
	return fmt.Sprintf("https://data.nba.net/data/10s/prod/v1/%s/scoreboard.json", date)
}
