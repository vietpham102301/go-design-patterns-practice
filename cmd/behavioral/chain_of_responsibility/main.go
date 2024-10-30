package main

import "github.com/vietpham102301/go-design-patterns-practice/bihavioral_patterns/chain_of_responsibility"

func main() {
	cashier := &chain_of_responsibility.Cashier{}

	medicineStore := &chain_of_responsibility.MedicineStore{}
	medicineStore.SetNext(cashier)

	doctor := &chain_of_responsibility.Doctor{}
	doctor.SetNext(medicineStore)

	reception := &chain_of_responsibility.Reception{}
	reception.SetNext(doctor)

	patient := &chain_of_responsibility.Patient{
		Name: "vietph21",
	}

	reception.Execute(patient)
}
