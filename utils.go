package main

import (
	"fmt"
	"math/rand"
)

func generateVariables() []*Variable {
	variables := make([]*Variable, VariablesLength)

	for i := 0; i < len(variables); i++ {
		variables[i] = NewVariableRandom()
	}

	return variables
}

func generateString(max int) string {
	random := make([]byte, max)

	if _, err := rand.Read(random); err != nil {
		panic(random)
	}

	return fmt.Sprintf("R%X", random)
}

// TODO: Better things
func generateInt(max int) string {
	return fmt.Sprintf("%d", rand.Intn(max))
}

func generateBool() string {
	if rand.Int()%2 == 0 {
		return "true"
	}

	return "false"
}
