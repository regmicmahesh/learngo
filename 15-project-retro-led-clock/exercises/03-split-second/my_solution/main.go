// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for {
		screen.Clear()
		screen.MoveTopLeft()
		hour, minute, second, millisecond := getCurrentTime()
		clockPrinter(hour, minute, second, millisecond)
		time.Sleep(100 * time.Millisecond)
	}

}

func clockPrinter(hour, minute, second, millisecond int) {

	fmt.Print("\n\n")
	for i := 0; i <= 10; i++ {

		//printing hour
		firstDigitBit, secondDigitBit := twoBitParser(hour)
		fmt.Print((digits[firstDigitBit])[i])
		fmt.Print(" ")
		fmt.Print((digits[secondDigitBit])[i])
		fmt.Print(" ")

		printSeparator(i, true)

		//printing minute
		firstDigitBit, secondDigitBit = twoBitParser(minute)
		fmt.Print((digits[firstDigitBit])[i])
		fmt.Print(" ")
		fmt.Print((digits[secondDigitBit])[i])
		fmt.Print(" ")

		printSeparator(i, (second%10)%2 == 0)

		//printing second
		firstDigitBit, secondDigitBit = twoBitParser(second)
		fmt.Print((digits[firstDigitBit])[i])
		fmt.Print(" ")
		fmt.Print((digits[secondDigitBit])[i])
		fmt.Print(" ")

		printSeparator(i, false)
		//printing second
		firstDigitBit, secondDigitBit = twoBitParser(millisecond)
		fmt.Print((digits[firstDigitBit])[i])
		fmt.Print(" ")
		fmt.Print((digits[secondDigitBit])[i])
		fmt.Print(" ")

		//ending the line
		fmt.Println("")
	}

}

func twoBitParser(num int) (x, y int) {
	x = num / 10
	y = num % 10
	return
}

func printSeparator(idx int, isVisible bool) {
	if isVisible {
		fmt.Print((seperator[idx]))
	} else {
		fmt.Print(strings.Repeat(" ", 13))
	}
}

func getCurrentTime() (int, int, int, int) {
	currentTime := time.Now()
	return currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond() / 1e8

}