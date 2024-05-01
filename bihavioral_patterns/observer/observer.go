package observer

import "fmt"

// Subject interface
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

// Observer interface
type Observer interface {
	Update(message string)
}

// ConcreteSubject stores state and notifies observers about changes
type ConcreteSubject struct {
	observers []Observer
	message   string
}

// RegisterObserver adds an observer to the list
func (cs *ConcreteSubject) RegisterObserver(o Observer) {
	cs.observers = append(cs.observers, o)
}

// RemoveObserver removes an observer from the list
func (cs *ConcreteSubject) RemoveObserver(o Observer) {
	var indexToRemove int
	for i, observer := range cs.observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}
	cs.observers = append(cs.observers[:indexToRemove], cs.observers[indexToRemove+1:]...)
}

// NotifyObservers alerts every observer about the change
func (cs *ConcreteSubject) NotifyObservers() {
	for _, observer := range cs.observers {
		observer.Update(cs.message)
	}
}

// UpdateMessage changes the message and notifies observers
func (cs *ConcreteSubject) UpdateMessage(message string) {
	cs.message = message
	cs.NotifyObservers()
}

// ConcreteObserver is a concrete observer that reacts to updates
type ConcreteObserver struct {
	Id int
}

// Update allows an observer to receive updates from the subject
func (co *ConcreteObserver) Update(message string) {
	fmt.Printf("Observer %d: received message: %s\n", co.Id, message)
}
