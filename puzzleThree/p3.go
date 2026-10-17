package puzzleThree

import (
	"fmt"
	"strings"
	"unicode"

	"stephenstucky.quest/flipflop/utils"
)

//https://flipflop.slome.org/2026/3
// Calculate the strength of passwords

// PartOne if the password contains at least one lowercase letter, add 1 to the score
// if the password contains at least one uppercase letter, add 1 to the score
// if there is at least one digit, add 1 to the score
// multiply this score by the length of the password
// find the strongest
func PartOne(input string) {

	passwords, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	strongest := ""
	highest := 0
	for _, password := range passwords {
		l := len(password)
		score := 0

		if hasLower(password) {
			score++
		}
		if hasUpper(password) {
			score++
		}
		if hasDigit(password) {
			score++
		}

		if score*l > highest {
			strongest = password
			highest = score * l
		}
	}
	fmt.Println("Strongest password:", strongest)
}

// PartTwo same rules as part 1 but some additional ones:
// - if the password contains the digit 7, and no other digits besides 7, add 7 to the score
// - if the password contains a sequence where the same character appears three or more times concsecutively, add the
// add the square of the length of the longest sequence to the score
// - if the password contains the lowercase string 'red', 'green', or 'blue' multiply the final score by 3 (applied once)
// multiply the resulting score by the length of the word.
func PartTwo(input string) {
	passwords, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	strongest := ""
	highest := 0
	for _, password := range passwords {
		l := len(password)
		score := 0

		if n := longestRun(password); n >= 3 {
			score += n * n
		}

		if onlySevens(password) {
			score += 7
		}

		if hasColor(password) {
			score *= 3
		}

		if score*l > highest {
			strongest = password
			highest = score * l
		}
	}

	fmt.Println("Strongest password:", strongest)
}

// PartThree append a single distinct character to each password, pick whatever character makes them as strong as possible
// this character will be appended to every password in the database
// any lowercase letter, uppercase letter, or digit is applicable
// all scoring rules from part 1 and part 2 apply
// find the sum of all the scores off passwords
func PartThree(input string) {
	passwords, err := utils.ReadStrings(input)
	if err != nil {
		panic(err)
	}

	candidates := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	bestTotal := 0
	bestChar := ' '
	for _, c := range candidates {
		total := 0
		for _, password := range passwords {
			total += calculateScore(password + string(c))
		}
		if total > bestTotal {
			bestTotal = total
			bestChar = c
		}
	}

	fmt.Printf("Best character: %c\n", bestChar)
	fmt.Println("Total Score:", bestTotal)
}

func hasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}

	return false
}

func hasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func hasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func onlySevens(s string) bool {
	has7 := false
	for _, r := range s {
		if r == '7' {
			has7 = true
		} else if unicode.IsDigit(r) {
			return false // some other digit
		}
	}
	return has7
}

func longestRun(s string) int {
	runes := []rune(s)
	longest, run := 0, 0
	for i, r := range runes {
		if i > 0 && r == runes[i-1] {
			run++
		} else {
			run = 1
		}
		if run > longest {
			longest = run
		}
	}
	return longest
}

func hasColor(s string) bool {
	for _, c := range []string{"red", "green", "blue"} {
		if strings.Contains(s, c) {
			return true
		}
	}
	return false
}

func calculateScore(s string) int {
	score := 0

	if hasLower(s) {
		score++
	}
	if hasUpper(s) {
		score++
	}
	if hasDigit(s) {
		score++
	}

	if onlySevens(s) {
		score += 7
	}
	if n := longestRun(s); n >= 3 {
		score += n * n
	}
	if hasColor(s) {
		score *= 3
	}

	return len(s) * score
}
