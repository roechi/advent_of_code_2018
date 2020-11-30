package main

import (
	"bufio"
	"fmt"
	"github.com/roechi/frequency"
	"log"
	"os"
	"strconv"
)

func readLines(filepath string) []string {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(file)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())

	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done...")

	return lines
}

func main() {
	lines := readLines("../puzzle_input.txt")

	ints := make([]int, len(lines))

	for i, s := range lines {
		ints[i], _ = strconv.Atoi(s)
	}

	var resultingFrequency = frequency.ApplyFrequencyChanges(ints, 0)
	fmt.Println("The resulting frequency is: ", resultingFrequency)

	duplicateFrequency := frequency.FindFirstDuplicateFrequency(ints, 0)
	fmt.Println("The first duplicated frequency is: ", duplicateFrequency)

}