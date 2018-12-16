package main

import (
	"fmt"
	"math/rand"
)

const (
	IntVariable int = 5
	MaxIntValue int = 1000
)

var (
	variableTypes = [...]string{"int"}
)

func generateString(max int) string {
	random := make([]byte, max)

	if _, err := rand.Read(random); err != nil {
		panic(random)
	}

	return fmt.Sprintf("R%X", random)
}

func variableGenerator() *Variable {
	switch variableTypes[rand.Intn(len(variableTypes))] {
	case "int":
		return NewVariable(rand.Intn(MaxIntValue), generateString(10), "int")
	}
	return NewVariable(0, generateString(10), "int")
}

func main() {
	var junkCode []*Variable

	for i := 0; i < IntVariable; i++ {
		junkCode = append(junkCode, variableGenerator())
	}

	fmt.Println(junkCode)

	Inject(junkCode)
}
