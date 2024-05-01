package main

import adapterPkg "github.com/vietpham102301/go-design-patterns-practice/structural_patterns/adapter"

func main() {
	// Our old system expects distances in miles, but the new system uses kilometers.
	adaptee := &adapterPkg.Adaptee{DistanceInKilometers: 100} // 100 kilometers
	adapter := &adapterPkg.Adapter{Adaptee: adaptee}

	// Client code can continue using the distance in miles without needing to know the details.
	adapterPkg.Client(adapter)
}
