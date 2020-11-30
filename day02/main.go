package main

import (
	"../utils"
	"./idscanner"
	"fmt"
)

func main() {
	ids := utils.ReadLines("./day02/puzzle_input.txt")

	checksum := idscanner.CalcChecksum(ids)

	fmt.Println("Checksum: ", checksum)

	equalChars := idscanner.FindEqualCharsOfMostSimilarId(ids)

	fmt.Println("Equal chars of two closest ids: ", equalChars)
}
