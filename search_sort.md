Certainly! Here are two well-known **sorting** algorithms and two well-known **searching** algorithms in **Go**, along with their time complexities.

---

### **1. Sorting Algorithms**

#### **Bubble Sort**
**Description**: Bubble sort is a simple comparison-based sorting algorithm. It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.

**Time Complexity**:
- **Best case**: O(n) when the list is already sorted.
- **Average case**: O(n²).
- **Worst case**: O(n²).

```go
package main

import "fmt"

// Bubble Sort Algorithm
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)  // Output: Sorted array: [11 12 22 25 64]
}
```

---

#### **Quick Sort**
**Description**: Quick sort is a divide-and-conquer algorithm. It selects a pivot element and partitions the array into two sub-arrays, one with elements smaller than the pivot and the other with elements greater than the pivot.

**Time Complexity**:
- **Best case**: O(n log n).
- **Average case**: O(n log n).
- **Worst case**: O(n²) (occurs when the pivot selection is poor, like in an already sorted array).

```go
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
	fmt.Println("Sorted array:", arr)  // Output: Sorted array: [11 12 22 25 64]
}
```

---

### **2. Searching Algorithms**

#### **Linear Search**
**Description**: Linear search is the simplest search algorithm. It checks every element in the list sequentially until it finds the target or reaches the end of the list.

**Time Complexity**:
- **Best case**: O(1) when the target element is at the first position.
- **Average case**: O(n).
- **Worst case**: O(n) when the target element is not in the list.

```go
package main

import "fmt"

// Linear Search Algorithm
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1 // return -1 if not found
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	target := 22
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Println("Element found at index:", index) // Output: Element found at index: 3
	} else {
		fmt.Println("Element not found")
	}
}
```

---

#### **Binary Search**
**Description**: Binary search is an efficient searching algorithm for sorted arrays. It repeatedly divides the search interval in half. If the value of the target is less than the value in the middle of the interval, the search continues on the left side, otherwise, it continues on the right.

**Time Complexity**:
- **Best case**: O(1) when the target is at the middle.
- **Average case**: O(log n).
- **Worst case**: O(log n).

```go
package main

import "fmt"

// Binary Search Algorithm
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid // target found
		} else if arr[mid] < target {
			left = mid + 1 // search right half
		} else {
			right = mid - 1 // search left half
		}
	}
	return -1 // target not found
}

func main() {
	arr := []int{11, 12, 22, 25, 64}
	target := 22
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Println("Element found at index:", index) // Output: Element found at index: 2
	} else {
		fmt.Println("Element not found")
	}
}
```

---

### **Time Complexity Summary:**

| Algorithm       | Best Case       | Average Case   | Worst Case      |
|-----------------|-----------------|----------------|-----------------|
| **Bubble Sort** | O(n)            | O(n²)          | O(n²)           |
| **Quick Sort**  | O(n log n)      | O(n log n)     | O(n²)           |
| **Linear Search** | O(1)           | O(n)           | O(n)            |
| **Binary Search** | O(1)           | O(log n)       | O(log n)        |

---

### Conclusion:
- **Sorting Algorithms**: **Bubble Sort** is easy but inefficient for large datasets. **Quick Sort** is more efficient with average-case time complexity of **O(n log n)**.
- **Searching Algorithms**: **Linear Search** is simple but inefficient, while **Binary Search** is much more efficient but requires a sorted array.

Each of these algorithms has its place depending on the use case. For example, if the data is already sorted or will be sorted once, binary search is a great option for fast lookups. Similarly, **Quick Sort** is generally a preferred sorting method for large datasets due to its efficiency.