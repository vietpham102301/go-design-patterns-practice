package main

import (
	"fmt"
	iteratorPkg "github.com/vietpham102301/go-design-patterns-practice/bihavioral_patterns/iterator"
)

func main() {
	list := iteratorPkg.NewLinkedList()
	list.Append("Hello")
	list.Append("World")
	list.Append(2024)

	iterator := list.Iterator()

	for iterator.HasNext() {
		value := iterator.Next()
		fmt.Println(value)
	}
}
