package main

import (
	"fmt"
	"math/rand"
)

const (
	VariablesLength int = 5
	MaxIntValue     int = 1000
	MaxStringLength int = 30
)

var (
	variableTypes = [...]string{"int", "string", "bool"}
)

type Variable struct {
	Value string
	Name  string
	Type  string
}

func NewVariable(v, n, t string) *Variable {
	return &Variable{
		Value: v,
		Name:  n,
		Type:  t,
	}
}

func NewVariableRandom() *Variable {
	switch variableTypes[rand.Intn(len(variableTypes))] {
	case "int":
		return NewVariable(generateInt(MaxIntValue), generateString(10), "int")
	case "string":
		return NewVariable(generateString(MaxStringLength), generateString(10), "string")
	//bool
	default:
		return NewVariable(generateBool(), generateString(10), "bool")
	}
}

// TODO
func (v *Variable) String() string {
	if v.Type == "string" {
		return fmt.Sprintf("var %s %s = \"%s\"", v.Name, v.Type, v.Value)
	}

	return fmt.Sprintf("var %s %s = %s", v.Name, v.Type, v.Value)
}
