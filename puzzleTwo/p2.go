package puzzleTwo

import (
	"fmt"
	"slices"

	"stephenstucky.quest/flipflop/utils"
)

// https://flipflop.slome.org/2026/2

const left = '<'
const right = '>'
const leftD = Direction("left")
const rightD = Direction("right")

// PartOne make sure the lazer doesn't overheat a part of the wall by firing at it too many times.
// the laser is controlled by > and < moving left and right. with wrap around.
// the wall has segments numbered 1 to 100
func PartOne(input string) {
	wall := make(map[int]int)
	position := 1

	inputData, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	commands := []rune(inputData[0])
	for _, cmd := range commands {
		if cmd == left {
			if position == 1 {
				position = 100
			} else {
				position = position - 1
			}
		}

		if cmd == right {
			if position == 100 {
				position = 1
			} else {
				position = position + 1
			}
		}

		wall[position]++ //safe because the default is 0
	}

	highestTempSegment := findHighestTempSegment(wall)
	ans := highestTempSegment * wall[highestTempSegment]
	fmt.Println("Answer:", ans)
}

func findHighestTempSegment(m map[int]int) int {
	highest := 0
	segment := 0

	for key, value := range m {
		// map iteration order is random, so ties must be broken explicitly
		if value > highest || (value == highest && key < segment) {
			highest = value
			segment = key
		}
	}

	return segment
}

// PartTwo there is now a movable wall segment. It executes commands in reverse order of the lazer.
// count how many times the movable wall segment is hit by the lazer
func PartTwo(input string) {
	lazerPos := 1
	wallPos := 1

	wallDegrees := 0

	commandInputs, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	lazerCommands := []rune(commandInputs[0])
	wallCommands := []rune(commandInputs[0])
	slices.Reverse(wallCommands)

	for i, lazerCmd := range lazerCommands {
		wallCmd := wallCommands[i]

		if lazerCmd == left {
			if lazerPos == 1 {
				lazerPos = 100
			} else {
				lazerPos = lazerPos - 1
			}
		}

		if lazerCmd == right {
			if lazerPos == 100 {
				lazerPos = 1
			} else {
				lazerPos = lazerPos + 1
			}
		}

		if wallCmd == left {
			if wallPos == 1 {
				wallPos = 100
			} else {
				wallPos = wallPos - 1
			}
		}

		if wallCmd == right {
			if wallPos == 100 {
				wallPos = 1
			} else {
				wallPos = wallPos + 1
			}
		}

		if lazerPos == wallPos {
			wallDegrees = wallDegrees + 1
		}
	}

	fmt.Println("Answer: ", wallDegrees)

}

// PartThree every wall segment actually moves around. because of course they do. they still execute commands in reverse order.
// every wall segment moves simultaneously on each command.
// find the hottest wall segment and multiply its _original_ segment number by its temp.
func PartThree(input string) {
	lazerPos := 1
	wall := make(map[int]*wallSegment, 100)
	for i := 1; i <= 100; i++ {
		wall[i] = &wallSegment{
			segmentNumber: i,
			temp:          0,
			currentPos:    i,
		}
	}

	commandInputs, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	lazerCommands := []rune(commandInputs[0])
	wallCommands := []rune(commandInputs[0])
	slices.Reverse(wallCommands)

	for i, lazerCmd := range lazerCommands {
		wallCmd := wallCommands[i]

		if lazerCmd == left {
			if lazerPos == 1 {
				lazerPos = 100
			} else {
				lazerPos = lazerPos - 1
			}
		}

		if lazerCmd == right {
			if lazerPos == 100 {
				lazerPos = 1
			} else {
				lazerPos = lazerPos + 1
			}
		}

		if wallCmd == left {
			moveAllWalls(leftD, wall)
		} else {
			moveAllWalls(rightD, wall)
		}
		findAndHeatWall(lazerPos, wall)
	}
	hottest := findHighestTempAgain(wall)
	if hottest == nil {
		panic("No hottest wall found!")
	}
	fmt.Println("Answer:", hottest.segmentNumber*hottest.temp)
}

type Direction string

type wallSegment struct {
	segmentNumber int
	temp          int
	currentPos    int
}

func (w *wallSegment) moveLeft() {
	if w.currentPos == 1 {
		w.currentPos = 100
	} else {
		w.currentPos = w.currentPos - 1
	}
}

func (w *wallSegment) moveRight() {
	if w.currentPos == 100 {
		w.currentPos = 1
	} else {
		w.currentPos = w.currentPos + 1
	}
}

func moveAllWalls(d Direction, walls map[int]*wallSegment) {
	for _, wall := range walls {
		if d == rightD {
			wall.moveRight()
		} else {
			wall.moveLeft()
		}
	}
}

func findAndHeatWall(lazerPos int, walls map[int]*wallSegment) {
	for _, wall := range walls {
		if wall.currentPos == lazerPos {
			wall.temp = wall.temp + 1
		}
	}
}

func findHighestTempAgain(m map[int]*wallSegment) *wallSegment {
	var hottest *wallSegment

	for _, seg := range m {
		// map iteration order is random, so ties must be broken explicitly
		if hottest == nil || seg.temp > hottest.temp ||
			(seg.temp == hottest.temp && seg.segmentNumber < hottest.segmentNumber) {
			hottest = seg
		}
	}

	return hottest
}
