package main

import (
	"fmt"
)

func maxScore(s string) int {
	ones, zeroes, score := 0, 0, 0

	// Find total ones in string
	for _, char := range s {
		if char == '1' {
			ones++
		}
	}

	// Traverse string

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroes++
		} else {
			ones--
		}
		score = max(score, zeroes+ones)
	}

	return score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// test cases
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
	fmt.Println(maxScore("1111"))
}
