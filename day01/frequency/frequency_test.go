package frequency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ApplyFrequencyChanges(t *testing.T) {
	want := 3
	changes := []int{1, 1, 1}
	got := ApplyFrequencyChanges(changes, 0)

	assert.Equal(t, want, got)
}

func Test_FindFirstDuplicateFrequency(t *testing.T) {
	want := 10
	changes := []int{3, 3, 4, -2, -4}

	got := FindFirstDuplicateFrequency(changes, 0)

	assert.Equal(t, want, got)
}
