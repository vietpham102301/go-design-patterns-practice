package factory_method

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK47",
			power: 100,
		},
	}
}
