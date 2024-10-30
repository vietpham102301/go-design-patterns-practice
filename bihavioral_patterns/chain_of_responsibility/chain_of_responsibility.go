package chain_of_responsibility

import "fmt"

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Patient struct {
	Name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Registration is already Done!")
		r.next.Execute(p)
		return
	}

	fmt.Println("Registration is Done!")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(department Department) {
	r.next = department
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup is already Done!")
		d.next.Execute(p)

		return
	}

	fmt.Println("Doctor checkup is Done!")
	d.next.Execute(p)
}

func (d *Doctor) SetNext(department Department) {
	d.next = department
}

type MedicineStore struct {
	next Department
}

func (m *MedicineStore) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine is already received!")
		m.next.Execute(p)

		return
	}

	fmt.Println("Medicine is received!")
	m.next.Execute(p)
}

func (m *MedicineStore) SetNext(department Department) {
	m.next = department
}

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Cashier already received payment from patient!")
	}
	fmt.Println("Cashier getting payment from patient!")
}

func (c *Cashier) SetNext(department Department) {
	c.next = department
}
