package factory_method

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 50,
		},
	}
}
