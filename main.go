package main

import "fmt"

func main() {
	fmt.Printf("Binary Search Algorithm\n")
	result := binarySearch([]int{2, 3, 4, 5, 56, 67, 78, 89, 123, 432, 567, 7654, 23456}, 78)
	fmt.Printf("Result: %v\n", result)
}

// Binary Search is a simple algorithm to find an item in an sorted array.
// Time: O(log n) - So very efficient
// The Go Standard Library already provides a binary tree implementation in it's sort package.
// sort.Search() - https://golang.org/pkg/sort/#Search
func binarySearch(list []int, search int) bool {
	startIndex := 0
	endIndex := len(list) - 1

	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2

		if list[median] < search {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}
	}

	if startIndex == len(list) || list[startIndex] != search {
		return false
	}
	return true
}
