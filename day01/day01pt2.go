package main

import (
	"strconv"
)

func rotationPt2(command string, dial int) int {
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
			if currDial-1 == 0 {
				erg++
			}
			if (currDial - 1) >= 0 {
				currDial -= 1

			} else {
				// erg++
				currDial = 99
			}
		}
	case 'R':
		for i := 0; i < turnInt; i++ {
			if (currDial + 1) <= 99 {
				currDial += 1
			} else {
				currDial = 0
				erg++
			}
		}
	}
	// if currDial == 0 {
	// 	erg++
	// }
	return currDial
}
