package frequency

import "github.com/roechi/set"

func ApplyFrequencyChanges(changes []int, baseFrequency int) int {
	for _, c := range changes {
		baseFrequency += c
	}

	return baseFrequency
}



func FindFirstDuplicateFrequency(changes []int, baseFrequency int) int {
	s := set.NewSet()
	idx := 0
	currentFrequency := baseFrequency

	for {
		if s.Contains(currentFrequency) {
			break
		} else {
			s.Add(currentFrequency)
			currentFrequency += changes[idx]
			idx++
			if len(changes) == idx {
				idx = 0
			}
		}
	}

	return currentFrequency
}



