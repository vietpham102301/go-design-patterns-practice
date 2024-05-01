package main

import "github.com/vietpham102301/go-design-patterns-practice/bihavioral_patterns/strategy"

func main() {
	sample := []int{64, 34, 25, 12, 22, 11, 90}

	// Using context to change strategy
	context := &strategy.Context{}

	// Apply Bubble sort
	context.SetStrategy(&strategy.BubbleSortStrategy{})
	context.Sort(append([]int(nil), sample...)) // Create a copy of the sample for sorting

	// Apply Quick sort
	context.SetStrategy(&strategy.QuickSortStrategy{})
	context.Sort(append([]int(nil), sample...)) // Create a copy of the sample for sorting

	// Apply Insertion sort
	context.SetStrategy(&strategy.InsertionSortStrategy{})
	context.Sort(append([]int(nil), sample...)) // Create a copy of the sample for sorting
}
