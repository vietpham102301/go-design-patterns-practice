package main

import (
	"fmt"
	builderPkg "github.com/vietpham102301/go-design-patterns-practice/creational_patterns/builder"
)

func main() {
	builder := builderPkg.NewCarBuilder()
	car := builder.SetMake("Toyota").
		SetModel("Corolla").
		SetColor("Red").
		SetYear(2021).Build()

	fmt.Printf("Car Built: %+v\n", car)
}
