package idscanner

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcChecksum(t *testing.T) {
	ids := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	want := 12

	got := CalcChecksum(ids)

	assert.Equal(t, want, got)
}

func TestCountOccurrences(t *testing.T) {
	given := "abcdef"
	want := []int{0, 0}

	got := CountOccurrences(given)

	assert.Equal(t, want, got)
}

func TestCountOccurrencesOneDoubleOneTriple(t *testing.T) {
	given := "bababc"
	want := []int{1, 1}

	got := CountOccurrences(given)

	assert.Equal(t, want, got)
}

func TestCountOccurrencesOneDouble(t *testing.T) {
	given := "abbcde"
	want := []int{1, 0}

	got := CountOccurrences(given)

	assert.Equal(t, want, got)
}

func TestCountOccurrencesOneTriple(t *testing.T) {
	given := "abcccd"
	want := []int{0, 1}

	got := CountOccurrences(given)

	assert.Equal(t, want, got)
}

func TestCountOccurrencesTwoDouble(t *testing.T) {
	given := "aabcdd"
	want := []int{1, 0}

	got := CountOccurrences(given)

	assert.Equal(t, want, got)
}

func TestFindEqualCharsOfMostSimilarId(t *testing.T) {
	given := []string{"fghij", "fguij"}
	want := "fgij"

	got := FindEqualCharsOfMostSimilarId(given)

	assert.Equal(t, want, got)
}

func TestFindEqualCharsOfMostSimilarIdForBatch(t *testing.T) {
	given := []string{"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz"}
	want := "fgij"

	got := FindEqualCharsOfMostSimilarId(given)

	assert.Equal(t, want, got)
}


