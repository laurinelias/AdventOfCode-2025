package main

import (
	"strconv"
	"strings"
)

func checkSequenzePt1(givenRange string) {
	tempRange := strings.Split(givenRange, "-")
	start, err := strconv.Atoi(string(tempRange[0]))
	end, err := strconv.Atoi(string(tempRange[1]))
	if err != nil {
		panic(err)
	}

	for i := start; i <= end; i++ {
		if start <= 99 {
			if checkOneDigitPt1(i) {
				sum += i
			}
		} else {
			if checkManyDigitsPt1(i) {
				sum += i
			}
		}
	}
}

func checkOneDigitPt1(oneDigit int) bool {
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

func checkManyDigitsPt1(manyDigit int) bool {
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
