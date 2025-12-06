package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func main() {

	p := Person{Name: "Sarvv"}
	fmt.Println(p.Speak())

}
