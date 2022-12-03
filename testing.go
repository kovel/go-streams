package main

import (
	"fmt"
)

type Pet struct {
	Type string
}

func (p Pet) String() string {
	return p.Type
}

type Human struct {
	Name string
	Pets *Vector[Pet]
}

func (h Human) String() string {
	return fmt.Sprintf("%s: [%v]", h.Name, h.Pets)
}

type Int int

func (i Int) String() string {
	return fmt.Sprint(int(i))
}
