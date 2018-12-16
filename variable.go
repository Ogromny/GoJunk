package main

import "fmt"

type Variable struct {
	Value int
	Name  string
	Type  string
}

func NewVariable(v int, n, t string) *Variable {
	return &Variable{
		Value: v,
		Name:  n,
		Type:  t,
	}
}

func (v *Variable) String() string {
	return fmt.Sprintf("var %s %s = %d", v.Name, v.Type, v.Value)
}
