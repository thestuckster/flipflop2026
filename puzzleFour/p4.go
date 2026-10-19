package puzzleFour

import (
	"fmt"
	"strings"

	"stephenstucky.quest/flipflop/utils"
)

// https://flipflop.slome.org/2026/4
// first 3 lines will always be the top of the flower
// the last line will always be soil
// leaves are either `o-|` or `|-o`

// PartOne how many leaves are harvested if we cut at line 400
// count the leaves above the cut
// we start counting at the first stem (so 1 is the first section of the stalk)
func PartOne(input string, cutAt int) {
	flower, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	//trim the first 3 lines, we know that's the top of the flower
	flower = flower[3:]

	//trim the last line, we know its the soil
	flower = flower[:len(flower)-1]

	leaves := 0
	for i := 0; i < len(flower)-cutAt; i++ {
		stalk := flower[i]
		if strings.Contains(stalk, "o-|") || strings.Contains(stalk, "|-o") {
			leaves++
		}
	}

	fmt.Println("Leaves:", leaves)
}

// PartTwo workers have to climb the stalk. but every time they encounter a leave on the opposite side, they have to swap
// how many times do the workers have to swap sides to get to the top of the stalk?
func PartTwo(input string) {
	flower, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	//trim the very top of the flower, it can't be climbed.
	flower = flower[3:]

	last := ""
	swaps := 0
	for _, stalk := range flower {
		if strings.Contains(stalk, "#") {
			//this is the ground. climb up.
			continue
		}

		if last == "" {
			if isLeftLeaf(stalk) {
				last = "left"
			} else {
				last = "right"
			}
			//first leaf we've seen. not a swap.
			continue
		}

		if isLeftLeaf(stalk) && last == "right" {
			swaps++
			last = "left"
		} else if isRightLeaf(stalk) && last == "left" {
			swaps++
			last = "right"
		}
	}

	fmt.Println("Swaps:", swaps)
}

func isLeftLeaf(stalk string) bool {
	return strings.Contains(stalk, "o-|")
}

func isRightLeaf(stalk string) bool {
	return strings.Contains(stalk, "|-o")
}

func PartThree(input string) {
	flower, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	//trim the actual flower, it can't be climbed
	flower = flower[3:]

	//walk the tree and get a list of all the leaves
	leaves := make([]byte, 0)
	for i := len(flower) - 1; i >= 0; i-- {
		//bottom up (since the first line of our input is the top of the flower)
		stalk := flower[i]
		if strings.Contains(stalk, "#") {
			//ground
			continue
		}

		if isLeftLeaf(stalk) {
			leaves = append(leaves, 'L')
		} else if isRightLeaf(stalk) {
			leaves = append(leaves, 'R')
		}
	}

	//simulate climbs
	workers := 0
	for len(leaves) > 0 {
		workers++
		survivingLeaves := leaves[:0] //new empty slice based on leaves.
		for i, side := range leaves {
			isTopLeaf := i == len(leaves)-1
			sameSideAbove := !isTopLeaf && leaves[i+1] == side
			if sameSideAbove {
				survivingLeaves = append(survivingLeaves, side)
			}
		}
		leaves = survivingLeaves
	}

	fmt.Println("Workers that can climb safely:", workers)
}
