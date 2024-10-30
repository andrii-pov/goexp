package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // v is available here in else block too
	}

	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switch os := runtime.GOOS; os { // switch statements in go always have implicit break
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s. \n", os)
	}

	today := time.Now().Weekday()
	fmt.Println("Today is", today, ", so Saturday is: ")
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case calcDayDiff(today, 1):
		fmt.Println("Tomorrow.")
	case calcDayDiff(today, 2):
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	fmt.Println("Based on your time, I'm greeting you: ")
	switch { // switch without a condition is the same as switch true, conditions are in cases
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func calcDayDiff(today time.Weekday, num int) time.Weekday {
	return today + time.Weekday(num)
}
