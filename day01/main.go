package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePt1, err := os.Open("day01pt1Input.txt")
	if err != nil {
		panic(err)
	}

	scannerPt1 := bufio.NewScanner(filePt1)
	for scannerPt1.Scan() {
		line := scannerPt1.Text()
		dial = rotationPt1(line, dial)
	}
	fmt.Println("NullCount Day 1Pt1:", erg)
	defer filePt1.Close()
	if err := scannerPt1.Err(); err != nil {
		panic(err)
	}

	dial = 50
	erg = 0
	filePt2, err := os.Open("day01pt2Input.txt")
	if err != nil {
		panic(err)
	}
	defer filePt2.Close()
	scannerPt2 := bufio.NewScanner(filePt2)
	for scannerPt2.Scan() {
		line := scannerPt2.Text()
		dial = rotationPt2(line, dial)
	}
	fmt.Println("NullCount Day 1Pt2:", erg)

	if err := scannerPt2.Err(); err != nil {
		panic(err)
	}

}
