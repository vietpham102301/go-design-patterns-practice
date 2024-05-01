package main

import (
	"fmt"
	"github.com/vietpham102301/go-design-patterns-practice/creational_patterns/singleton"
)

func main() {

	// goroutine safe.
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 30; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		obj := singleton.GetInstance()
	//		fmt.Println(obj.Value)
	//	}()
	//}
	//
	//wg.Wait() // Wait for all goroutines to finish
	//fmt.Println("All goroutines finished executing.")

	obj1 := singleton.GetInstance()

	obj2 := singleton.GetInstance()

	if obj1 == obj2 {
		fmt.Println("it's identical")
	}
}
