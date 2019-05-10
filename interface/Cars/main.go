package main

import (
	"fmt"
)

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type Says interface {
	Hello()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func (p *BMW) Hello() {
	fmt.Printf("Hello,I`m %s \n", p.Name)
}

func main() {
	var car Car
	fmt.Println(car)

	bmw := &BMW{
		Name: "BMW",
	}
	car = bmw
	car.Run()
	car.DiDi()
	fmt.Println(car.GetName())

	bmw.Hello()

}
