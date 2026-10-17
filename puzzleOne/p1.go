package puzzleOne

import (
	"fmt"
	"slices"

	"stephenstucky.quest/flipflop/utils"
)

// PartOne https://flipflop.slome.org/2026/1
// how many seconds does it take total to heat up all the cups of coffee to at least 60 degrees
func PartOne(inputPath string) {
	temps, err := utils.ReadInts(inputPath)

	if err != nil {
		panic(err)
	}

	totalSeconds := 0
	for _, temp := range temps {
		if temp < 60 {
			totalSeconds += 60 - temp
		}
	}

	fmt.Printf("Total seconds: %d\n", totalSeconds)
}

// PartTwo ensure all coffees are 60 degrees
// to heat up a cup it takes 1 second per degree
// to cool down a cup it takes 5 seocnds per degree
func PartTwo(inputPath string) {
	temps, err := utils.ReadInts(inputPath)
	if err != nil {
		panic(err)
	}

	totalSeconds := 0
	for _, temp := range temps {
		if temp < 60 {
			totalSeconds += 60 - temp
		}
		if temp > 60 {
			degrees := temp - 60
			totalSeconds += degrees * 5
		}
	}
	fmt.Printf("Total seconds: %d\n", totalSeconds)
}

// PartThree the input actually contains two pieces of information
// the first part is the current temp
// the second part is the desired temp
// go through the list and determine how many seconds it takes to reach the desired temp using the logic from previous parts
func PartThree(inputPath string) {

	totalSeconds := 0

	temps, err := utils.ReadInts(inputPath)
	if err != nil {
		panic(err)
	}

	current, desired := splitSlice(temps)
	for i, currentTemp := range current {
		desiredTemp := desired[i]

		if currentTemp < desiredTemp {
			totalSeconds += desiredTemp - currentTemp
		} else if currentTemp > desiredTemp {
			degrees := currentTemp - desiredTemp
			totalSeconds += degrees * 5
		}
	}
	fmt.Printf("Total seconds: %d\n", totalSeconds)
}

func splitSlice(s []int) (left, right []int) {
	m := len(s) / 2
	left = slices.Clone(s[:m])
	right = slices.Clone(s[m:])

	return left, right
}
