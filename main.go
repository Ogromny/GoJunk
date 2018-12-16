package main

import (
	"math/rand"
	"time"
)

const (
	// How many variables to generate
	NumberOfVariables int = 5

	// Maximum length for generated variable name
	MaximumLengthOfVariableName int = 30

	// Maximum int value for int and bool
	MaximumIntValue int = 1000

	// Maximum length of generated string
	MaximumStringLength int = 30

	// Bool Junk Complexity ( need to be odd ! )
	BoolComplexity int = 13
)

func main() {
	rand.Seed(time.Now().Unix())

	variables := make([]*Variable, NumberOfVariables)

	for i:= 0; i < len(variables); i++ {
		variables[i] = NewVariableRandom()
	}

	Inject(variables)
}
