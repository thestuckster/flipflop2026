package main

import (
	"fmt"

	"stephenstucky.quest/flipflop/puzzleFour"
	"stephenstucky.quest/flipflop/puzzleOne"
	"stephenstucky.quest/flipflop/puzzleThree"
	"stephenstucky.quest/flipflop/puzzleTwo"
)

func main() {

	fmt.Println("Part One")
	puzzleOne.PartOne("./puzzleOne/p1-1.txt")
	puzzleOne.PartTwo("./puzzleOne/p1-2.txt")
	puzzleOne.PartThree("./puzzleOne/p1-3.txt")

	fmt.Println("\nPart Two")
	puzzleTwo.PartOne("./puzzleTwo/input.txt")
	puzzleTwo.PartTwo("./puzzleTwo/input.txt")
	puzzleTwo.PartThree("./puzzleTwo/input.txt")

	fmt.Println("\nPart Three")
	puzzleThree.PartOne("./puzzleThree/input.txt")
	puzzleThree.PartTwo("./puzzleThree/input.txt")
	puzzleThree.PartThree("./puzzleThree/input.txt")

	fmt.Println("\nPart Four")
	puzzleFour.PartOne("./puzzleFour/input.txt", 400)

}
