package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePt1, err := os.Open("day02pt1Input.txt")
	if err != nil {
		panic(err)
	}
	defer filePt1.Close()
	scannerPt1 := bufio.NewScanner(filePt1)
	for scannerPt1.Scan() {
		tempInput := strings.Split(scannerPt1.Text(), ",")
		for _, element := range tempInput {
			checkSequenzePt1(element)
		}
		fmt.Println("sum", sum)
	}

	sum = 0
	filePt2, err := os.Open("day02pt2Input.txt")
	if err != nil {
		panic(err)
	}
	defer filePt2.Close()
	scannerPt2 := bufio.NewScanner(filePt2)
	for scannerPt2.Scan() {
		tempInput := strings.Split(scannerPt2.Text(), ",")
		for _, element := range tempInput {
			checkSequenzePt2(element)
		}
		fmt.Println("sum", sum)
	}
}
