package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeTime(pmHour, minute int, location, date string) time.Time {
	l, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, time.April, 2, 12+pmHour, 0, 0, 0, l) // known cavs game
	return t
}

type FakeClock struct {
	now time.Time
}

func newFakeClock(startTime time.Time) *FakeClock {
	return &FakeClock{startTime}
}

func (c *FakeClock) Now() time.Time {
	return c.now
}

func (c *FakeClock) Advance(d time.Duration) {
	c.now = c.now.Add(d)
}

type fakeScheduleURL struct {
	url string
}

func newFakeScheduleURL(url string) fakeScheduleURL {
	return fakeScheduleURL{url: url}
}

func (r fakeScheduleURL) URL() string {
	return r.url
}

func newFixtureHandlerFunc(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
		fmt.Fprintf(w, "%s", contents)
	}
}
