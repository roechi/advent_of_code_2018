package main

import (
	"../utils"
	"./frequency"
	"fmt"
	"strconv"
)

func main() {
	lines := utils.ReadLines("./day01/puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	var resultingFrequency = frequency.ApplyFrequencyChanges(ints, 0)
	fmt.Println("The resulting frequency is: ", resultingFrequency)

	duplicateFrequency := frequency.FindFirstDuplicateFrequency(ints, 0)
	fmt.Println("The first duplicated frequency is: ", duplicateFrequency)

}
