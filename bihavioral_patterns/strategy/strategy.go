package strategy

import (
	"fmt"
	"sort"
)

// SortStrategy defines the interface for sorting strategies
type SortStrategy interface {
	Sort([]int)
}

// Context defines the environment in which strategies are used
type Context struct {
	strategy SortStrategy
}

// SetStrategy allows the client to change the strategy
func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

// Sort uses the current strategy to sort the array
func (c *Context) Sort(array []int) {
	c.strategy.Sort(array)
}

// BubbleSortStrategy implements bubble sort
type BubbleSortStrategy struct{}

func (s *BubbleSortStrategy) Sort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("Sorted using bubble sort:", array)
}

// QuickSortStrategy implements quick sort using sort package
type QuickSortStrategy struct{}

func (s *QuickSortStrategy) Sort(array []int) {
	sort.Ints(array)
	fmt.Println("Sorted using quick sort:", array)
}

// InsertionSortStrategy implements insertion sort
type InsertionSortStrategy struct{}

func (s *InsertionSortStrategy) Sort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1

		// Move elements of array[0..i-1], that are
		// greater than key, to one position ahead
		// of their current position
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
	fmt.Println("Sorted using insertion sort:", array)
}
