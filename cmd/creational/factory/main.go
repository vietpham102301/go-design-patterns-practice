package main

import (
	"fmt"
	"github.com/vietpham102301/go-design-patterns-practice/creational_patterns/factory_method"
)

func main() {
	ak47, _ := factory_method.GetGun("AK47")
	musket, _ := factory_method.GetGun("Musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(gun factory_method.IGun) {
	fmt.Printf("name: %s\n", gun.GetName())
	fmt.Printf("power: %d\n", gun.GetPower())
}
