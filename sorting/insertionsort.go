package main

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		fmt.Println("key", key)

		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
			fmt.Println("arr", arr)

		}
		arr[j+1] = key
		fmt.Println("arr", arr, "key", key)
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	// quickSort(arr)
	// fmt.Println("Sorted array:", arr) // Output: Sorted array: [11 12 22 25 64]
	InsertionSort(arr)
	fmt.Println("Insertion Sorted ", arr)
}
