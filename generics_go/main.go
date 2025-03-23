package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 5, 10
	x, y = swap(x, y)
	fmt.Println(x, y)

	a, b := "a", "b"
	a, b = swap(a, b)
	fmt.Println(a, b)

}
