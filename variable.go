package main

import (
	"fmt"
	"math/rand"
)

var (
	variableTypes = [...]string{"int", "string", "bool"}
)

type Variable struct {
	Value string
	Name  string
	Type  string
}

func NewVariableRandom() *Variable {
	switch variableTypes[rand.Intn(len(variableTypes))] {
	case "int":
		return &Variable{
			Value: fmt.Sprintf("%d", randomInt(0, MaximumIntValue)),
			Name:  randomString(MaximumLengthOfVariableName),
			Type:  "int",
		}
	case "string":
		return &Variable{
			Value: randomString(MaximumStringLength),
			Name:  randomString(MaximumLengthOfVariableName),
			Type:  "string",
		}
	default:
		return &Variable{
			Value: randomBool(),
			Name:  randomString(MaximumLengthOfVariableName),
			Type:  "bool",
		}
	}
}

func (v *Variable) String() string {
	format := "var %s %s = "

	if v.Type == "string" {
		format += "\"%s\""
	} else {
		format += "%s"
	}

	return fmt.Sprintf(format, v.Name, v.Type, v.Value)
}
