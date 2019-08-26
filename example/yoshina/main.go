package main

import (
	"fmt"

	"go.uber.org/fx"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func Show(p *Person) {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

func main() {
	fx.New(
		fx.Provide(
			func() int { return 27 },
			NewPerson,
			func() string { return "taro" },
		),
		fx.Invoke(Show),
	)
}
