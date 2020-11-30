package idscanner

import (
	"strings"
)

func CalcChecksum(ids []string) int {
	sumCounts := []int{0, 0}
	for _, id := range ids {
		counts := CountOccurrences(id)
		sumCounts[0] += counts[0]
		sumCounts[1] += counts[1]
	}

	return sumCounts[0] * sumCounts[1]
}

func CountOccurrences(id string) []int {
	result := make([]int, 2)
	for _, c := range id {
		count := strings.Count(id, string(c))
		if count == 2 {
			result[0] = 1
		} else if count == 3 {
			result[1] = 1
		}
	}
	return result
}

func FindEqualCharsOfMostSimilarId(ids []string) string {
	pos := 0
	for pos < len(ids[0]) {
		m := make(map[string]bool)
		for _, id := range ids {
			var shortId string
			if pos == len(id) - 1 {
				shortId = id[:pos]
			} else {
				shortId = id[:pos] + id[pos + 1:]
			}
			if m[shortId] {
				return shortId
			} else {
				m[shortId] = true
			}
		}
		pos++
	}
	return ""
}
