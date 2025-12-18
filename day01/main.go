package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("day01pt1Input.txt")
	file, err := os.Open("day01pt2Input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// dial = rotationPt1(line, dial)
		dial = rotationPt2(line, dial)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Println("NullCount Day 1Pt1:", erg)
	fmt.Println("NullCount Day 1Pt2:", erg)
}
