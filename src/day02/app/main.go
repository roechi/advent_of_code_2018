package main

import (
	"fmt"
	"github.com/roechi/idscanner"
	"github.com/roechi/utils"
)

func main() {
	ids := utils.ReadLines("../puzzle_input.txt")

	checksum := idscanner.CalcChecksum(ids)

	fmt.Println("Checksum: ", checksum)

	equalChars := idscanner.FindEqualCharsOfMostSimilarId(ids)

	fmt.Println("Equal chars of two closest ids: ", equalChars)
}
