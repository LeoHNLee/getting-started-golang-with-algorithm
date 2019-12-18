package main

// About bufio pacakge
// https://golang.org/pkg/bufio
// About strings package
// https://golang.org/pkg/strings

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func no2753() {
	// Read the Text
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// convert text to int
	year, _ := strconv.Atoi(text)

	// compute: if-else is more confused
	switch {
	case year%400 == 0:
		fmt.Println(1)
	case year%100 == 0:
		fmt.Println(0)
	case year%4 == 0:
		fmt.Println(1)
	default:
		fmt.Println(0)
	}
}
