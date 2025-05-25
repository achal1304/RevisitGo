package main

import "fmt"

// Quick Sort Algorithm
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Choose pivot
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	middle := []int{}

	// Partitioning step
	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v > pivot:
			right = append(right, v)
		default:
			middle = append(middle, v)
		}
	}

	// Recursively sort left and right partitions
	quickSort(left)
	quickSort(right)

	// Concatenate everything back together
	copy(arr, append(append(left, middle...), right...))
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	quickSort(arr)
	fmt.Println("Sorted array:", arr) // Output: Sorted array: [11 12 22 25 64]
}
