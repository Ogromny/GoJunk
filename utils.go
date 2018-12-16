package main

import (
	"fmt"
	"math/rand"
)

func randomString(max int) string {
	random := make([]byte, max)

	if _, err := rand.Read(random); err != nil {
		panic(random)
	}

	return fmt.Sprintf("R%X", random[:len(random)-1])
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomBool() string {
	boolOperator := [...]string{"<", ">", "=="}
	boolSeparator := [...]string{"||", "&&"}

	var junk string

	for i := 0; i < BoolComplexity; i++ {
		random := randomInt(0, MaximumIntValue)

		junk += fmt.Sprintf("%d ", random)

		if i%2 == 0 {
			junk += fmt.Sprintf("%s ", boolOperator[random%len(boolOperator)])
		} else {
			junk += fmt.Sprintf("%s ", boolSeparator[random%len(boolSeparator)])
		}

		if i == BoolComplexity-1 {
			junk += fmt.Sprintf("%d", randomInt(0, MaximumIntValue))
		}
	}

	return fmt.Sprintf("(%s)", junk)
}
