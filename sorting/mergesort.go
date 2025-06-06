package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	fmt.Println(arr)

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	if len(left) == 0 && len(right) == 0 {
		return []int{}
	}

	l, r := 0, 0
	result := []int{}
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
		fmt.Println("left. righjt", left, right, l, r)
	}

	remaining := append(left[l:], right[r:]...)
	return append(result, remaining...)
}

// func main() {
// 	arr := []int{64, 25, 12, 22, 11}
// 	// quickSort(arr)
// 	// fmt.Println("Sorted array:", arr) // Output: Sorted array: [11 12 22 25 64]

// 	fmt.Println("Mergen Sorted ", MergeSort(arr))
// }
