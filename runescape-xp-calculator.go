package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//read arguments
	if len(os.Args) == 5 {
		currentXp, _ := strconv.ParseInt(os.Args[1], 10, 32)
		requestedLevel, _ := strconv.ParseInt(os.Args[2], 10, 32)
		xpPerAction, _ := strconv.ParseFloat(os.Args[3], 32)
		gpPerAction, _ := strconv.ParseInt(os.Args[4], 10, 32)

		//do calculations
		currentLevel := XpToLevel(currentXp)
		neededXp := int64(levelToXp(requestedLevel))
		xpDifference := neededXp - currentXp
		actionsNeeded := float64(xpDifference) / xpPerAction

		//print results

		fmt.Println("Current level:", currentLevel)
		fmt.Println("Requested level:", requestedLevel)
		fmt.Println("XP required:", xpDifference)
		fmt.Println("You need to train:", actionsNeeded, "times")
		fmt.Println("This will cost you:", float64(gpPerAction)*actionsNeeded, "gp")
	} else {
		//output error
		fmt.Println("usage: runescape-xp-calculator [currentXp] [requestedLevel] [xpPerAction] [gpPerAction]")
	}
}

func levelToXp(level int64) int {
	return int(levels[level-1])
}

func XpToLevel(xp int64) int {
	for level, xpForLevel := range levels {
		if xpForLevel > xp {
			return level
		}
	}
	return 0
}
