package t_bank

import (
	"math"
	"strings"
	"unicode"
)

const (
	target string = "rumpelstiltskin"
)

func countWord(s string) int {
	s = strings.ToLower(s)

	inputFreq := make(map[rune]int)
	for _, r := range s {
		if unicode.IsLetter(r) {
			inputFreq[r]++
		}
	}

	targetFreq := make(map[rune]int)
	for _, r := range target {
		targetFreq[r]++
	}

	minCount := math.MaxInt
	for ch, count := range targetFreq {
		available := inputFreq[ch]
		min := available / count
		if min < minCount {
			minCount = min
		}
	}

	return minCount
}
