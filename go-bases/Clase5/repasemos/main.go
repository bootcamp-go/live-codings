package main

import "fmt"

type Bootcamper struct {
	Name     string `json:"name"`
	Age      uint   `json:"age"`
	Position string `json:"position"`
}

func (b *Bootcamper) Print() {
	fmt.Println("===Bootcamper===")
	fmt.Printf("Name: %s\n", b.Name)
	fmt.Printf("Age: %d\n", b.Age)
	fmt.Printf("Position: %s\n", b.Position)
}

func (b *Bootcamper) Set(name string, age uint, position string) {
	b.Name = name
	b.Age = age
	b.Position = position
}

func main() {

	boot := Bootcamper{}

	boot.Set("Vincent", 35, "Backend Developer")
	boot.Print()

}
