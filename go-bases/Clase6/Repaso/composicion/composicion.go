package main

import (
	"fmt"
)

type animal struct {
	name    string
	species string
}

func (a *animal) Name() {
	fmt.Println(a.name)
}

func (a *animal) Species() {
	fmt.Printf("%s belongs to %s species\n", a.name, a.species)
}

type gopher struct {
	color string
	animal
}

func (g *gopher) Color() {
	fmt.Printf("%s is the color: %s\n", g.name, g.color)
}

func main() {
	a := animal{"Gopher", "Go Gopher"}
	g := &gopher{"blue", a}

	g.Name()
	g.Species()
	g.Color()
}
