package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var erg int

var dial int = 50

func main() {
	file, err := os.Open("day01pt1Input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dial = rotation(line, dial)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("NullCount", erg)
}

func rotation(command string, dial int) int {
	var currDial = dial
	direction := command[:1]
	turn := command[1:]
	turnInt, err := strconv.Atoi(string(turn))
	if err != nil {
		panic(err)
	}

	switch direction[0] {
	case 'L':
		for i := 0; i < turnInt; i++ {
			if (currDial - 1) >= 0 {
				currDial -= 1
			} else {
				currDial = 99
			}
		}
	case 'R':
		for i := 0; i < turnInt; i++ {
			if (currDial + 1) <= 99 {
				currDial += 1
			} else {
				currDial = 0
			}
		}
	}
	if currDial == 0 {
		erg++
	}
	return currDial
}
