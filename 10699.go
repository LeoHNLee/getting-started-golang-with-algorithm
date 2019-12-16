package main

// how to handle Time in Golang
// About "time" package
// https://golang.org/pkg/time/#hdr-Monotonic_Clocks

import (
	"fmt"
	"time"
)

func main() {
	// Set Time Zone of Seoul
	seoulTZ := time.FixedZone("Asia/Seoul", 9*60*60) // if you know utc+-hour or others, you can insert second
	seoulTZ, err := time.LoadLocation("Asia/Seoul")  // you know only TZ name
	if err != nil {
		panic(err)
	}

	// Get Now Seoul Time
	nowSeoul := time.Now().In(seoulTZ)

	// print out with format
	// golang reference date: 20060102150405
	fmt.Println(nowSeoul.Format("2006-01-02")) // YOU MUST USE ONLY THAT SPECIFIC YEAR-MONTH-DAY FORMAT
	fmt.Printf("%d-%02d-%02d", nowSeoul.Year(), int(nowSeoul.Month()), nowSeoul.Day())
}
