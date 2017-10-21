package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type fakeScheduleURL struct {
	url string
}

func newFakeScheduleURL(url string) fakeScheduleURL {
	return fakeScheduleURL{url: url}
}

func (r fakeScheduleURL) URL() string {
	return r.url
}

func newFixtureHandlerFunc() http.HandlerFunc {
	var boxscoreCounter = 1
	return func(w http.ResponseWriter, r *http.Request) {
		var filename string

		switch r.URL.Path {
		case "/prod/v1/20170609/scoreboard.json":
			filename = "fixtures/scoreboard.json"
		case "/prod/v1/20170609/0041600404_pbp_1.json":
			filename = "fixtures/pbp_1.json"
		case "/prod/v1/20170609/0041600404_pbp_2.json":
			filename = "fixtures/pbp_2.json"
		case "/prod/v1/20170609/0041600404_pbp_3.json":
			filename = "fixtures/pbp_3.json"
		case "/prod/v1/20170609/0041600404_pbp_4.json":
			filename = "fixtures/pbp_4.json"
		case "/prod/v1/20170609/0041600404_mini_boxscore.json":
			filename = fmt.Sprintf("fixtures/mini_boxscore_%d.json", boxscoreCounter)
			boxscoreCounter += 1
		}

		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
		fmt.Fprintf(w, "%s", contents)
	}
}
