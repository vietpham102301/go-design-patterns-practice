package adapter

import "fmt"

// Target interface represents the expected interface by the client
type Target interface {
	GetDistanceInMiles() float64
}

// Adaptee is the class that has the incompatible interface
type Adaptee struct {
	DistanceInKilometers float64
}

// GetDistance returns the distance in kilometers
func (a *Adaptee) GetDistanceInKilometers() float64 {
	return a.DistanceInKilometers
}

// Adapter makes Adaptee's interface compatible with the Target interface
type Adapter struct {
	*Adaptee
}

// GetDistanceInMiles converts the distance from kilometers to miles
func (adapter *Adapter) GetDistanceInMiles() float64 {
	return adapter.Adaptee.GetDistanceInKilometers() * 0.621371
}

// Client uses the Target interface
func Client(target Target) {
	fmt.Printf("Distance in miles: %.2f\n", target.GetDistanceInMiles())
}
