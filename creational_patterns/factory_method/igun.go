package factory_method

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
