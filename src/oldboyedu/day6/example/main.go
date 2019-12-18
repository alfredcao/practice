package main

import "fmt"

type Car interface {
	getName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) getName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running!\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi!\n", p.Name)
}

func main() {
	var car Car = &BMW{
		Name: "BMW",
	}

	car.Run()
	car.DiDi()
}
