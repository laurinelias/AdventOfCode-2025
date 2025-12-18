package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum int

func main() {
	file, err := os.Open("day02pt1Input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tempInput := strings.Split(scanner.Text(), ",")
		for _, element := range tempInput {
			checkSequenze(element)
		}
		fmt.Println("sum", sum)
	}
}

func checkSequenze(givenRange string) {
	tempRange := strings.Split(givenRange, "-")
	start, err := strconv.Atoi(string(tempRange[0]))
	end, err := strconv.Atoi(string(tempRange[1]))
	if err != nil {
		panic(err)
	}

	for i := start; i <= end; i++ {
		if start <= 99 {
			if checkOneDigit(i) {
				sum += i
			}
		} else {
			if checkManyDigits(i) {
				sum += i
			}
		}
	}
}

func checkOneDigit(oneDigit int) bool {
	s := strconv.Itoa(oneDigit)

	if oneDigit <= 99 {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				return true
			}
		}
		return false

	} else {
		return false
	}
}

func checkManyDigits(manyDigit int) bool {
	s := strconv.Itoa(manyDigit)
	if len(s)%2 == 0 {

		mid := len(s) / 2
		sequence := s[:mid]
		checkHalf := s[mid:]

		if sequence == checkHalf {
			return true
		}
	}
	return false
}
