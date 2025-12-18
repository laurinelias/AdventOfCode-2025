package main

import (
	"strconv"
)

func rotationPt1(command string, dial int) int {
	currDial := dial
	direction := command[:1]
	turn := command[1:]

	turnInt, err := strconv.Atoi(turn)
	if err != nil {
		panic(err)
	}

	switch direction[0] {
	case 'L':
		for i := 0; i < turnInt; i++ {
			if currDial > 0 {
				currDial--
			} else {
				currDial = 99
			}
		}
	case 'R':
		for i := 0; i < turnInt; i++ {
			if currDial < 99 {
				currDial++
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
