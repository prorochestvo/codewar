package codewar

import (
	"math"
	"strings"
)

// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272/train/go
func HighestScoringWord(s string) string {
	result := ""
	maxScore := math.MinInt64
	words := strings.Fields(s)
	for _, word := range words {
		score := 0
		for _, letter := range word {
			score += int(letter) - int('a') + 1
		}
		if score > maxScore {
			maxScore = score
			result = word
		}
	}
	return result
}
