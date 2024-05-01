package main

import "github.com/vietpham102301/go-design-patterns-practice/bihavioral_patterns/observer"

func main() {
	// Create a subject
	subject := &observer.ConcreteSubject{}

	// Create observers
	observer1 := &observer.ConcreteObserver{Id: 1}
	observer2 := &observer.ConcreteObserver{Id: 2}

	// Register observers
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	// Change subject and notify observers
	subject.UpdateMessage("Hello, World!")

	// Remove an observer
	subject.RemoveObserver(observer1)

	// Change subject and notify observers again
	subject.UpdateMessage("Second Message")
}
