package main

import (
	"fmt"
)

type dieRollFunc func(int) int

func fakeDieRoll(size int) int {
	return 42
}

func getDieRolls() []dieRollFunc{
	return []dieRollFunc{
		dieRoll,
		fakeDieRoll,
	}
}

var rolls = getDieRolls()
for index, rollFunc := range rolls{
	fmt.Printf("Die Roll Attempt #%d, result: %\n", index,rollFunc(10))
}
