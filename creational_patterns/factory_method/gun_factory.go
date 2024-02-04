package factory_method

import "fmt"

func GetGun(gunName string) (IGun, error) {
	if gunName == "AK47" {
		return newAK47(), nil
	}
	if gunName == "Musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
