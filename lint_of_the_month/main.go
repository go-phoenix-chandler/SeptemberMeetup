package main

import (
	"errors"
	"fmt"
)

type class int

const (
	mammal  class = 0
	insect        = 1
	bird          = 2
	reptile       = 3
)

type walks interface {
	walk()
}

type human struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	alive bool   `json:"alive"`
}

func newHuman(n string, a int, l bool) *human {
	return &human{
		Name:  n,
		Age:   a,
		alive: l,
	}
}

func (h *human) walk() {
	fmt.Printf("I'm walking...\n")
}

type animal struct {
	Legs int
}

func newAnimal(c class) (*animal, error) {
	if c == mammal {
		return &animal{4}, nil
	}
	if c == insect {
		return &animal{6}, nil
	}
	if c == bird {
		return &animal{2}, nil
	}
	if c == reptile {
		return &animal{4}, nil
	}
	return nil, errors.New("unknown animal class")
}

func (a *animal) walk() {
	fmt.Printf("I'm walking...\n")
}

type dog struct {
	Base *animal `json:"base"`
}

func newDog() *dog {
	a, _ := newAnimal(mammal)
	return &dog{
		Base: a,
	}
}

func main() {
	felix := newDog()
	george := newDog()

	felix.Base.walk()
	george.Base.walk()
}
