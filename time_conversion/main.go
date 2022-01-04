package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}

func timeConversion(s string) string {
	amPMLayout := "03:04:05PM"

	timeConverted, _ := time.Parse(amPMLayout, s)

	militaryLayout := "15:04:05"
	return timeConverted.Format(militaryLayout)
}
