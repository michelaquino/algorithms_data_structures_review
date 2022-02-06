package main

import (
	"fmt"
	"time"
)

type class struct {
	name  string
	start time.Time
	end   time.Time
}

func main() {
	classes := []class{
		{
			name:  "art",
			start: time.Date(2020, 01, 01, 9, 0, 0, 0, &time.Location{}),
			end:   time.Date(2020, 01, 01, 10, 0, 0, 0, &time.Location{}),
		},
		{
			name:  "english",
			start: time.Date(2020, 01, 01, 9, 30, 0, 0, &time.Location{}),
			end:   time.Date(2020, 01, 01, 10, 30, 0, 0, &time.Location{}),
		},
		{
			name:  "math",
			start: time.Date(2020, 01, 01, 10, 0, 0, 0, &time.Location{}),
			end:   time.Date(2020, 01, 01, 11, 0, 0, 0, &time.Location{}),
		},
		{
			name:  "computer_science",
			start: time.Date(2020, 01, 01, 10, 30, 0, 0, &time.Location{}),
			end:   time.Date(2020, 01, 01, 11, 30, 0, 0, &time.Location{}),
		},
		{
			name:  "music",
			start: time.Date(2020, 01, 01, 11, 0, 0, 0, &time.Location{}),
			end:   time.Date(2020, 01, 01, 12, 0, 0, 0, &time.Location{}),
		},
	}

	bestClasses := findBestSchedule(classes)
	for _, c := range bestClasses {
		fmt.Printf("%v - %v -> %s\n", c.start.Format(time.Kitchen), c.end.Format(time.Kitchen), c.name)
	}
}

func findBestSchedule(classes []class) map[string]class {
	schedule := map[string]class{}
	lastClass := class{}

	for {
		class := endEarly(classes, lastClass, schedule)
		if class.name == "" {
			break
		}

		schedule[class.name] = class
		lastClass = class
	}

	return schedule
}

func endEarly(classes []class, lastClass class, toDisconsider map[string]class) class {
	var earlyClass class
	earlyTime := time.Date(2100, 01, 01, 0, 0, 0, 0, &time.Location{})

	for _, c := range classes {
		if _, found := toDisconsider[c.name]; found {
			continue
		}

		if lastClass.name != "" && c.start.Before(lastClass.end) {
			continue
		}

		if c.end.Before(earlyTime) {
			earlyTime = c.end
			earlyClass = c
		}
	}

	return earlyClass
}
