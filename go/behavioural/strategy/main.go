package main

import "fmt"

func binarySearch(x int, xs []int) bool {

	low := 0
	high := len(xs) - 1

	for low <= high {
		median := (low + high) / 2

		if xs[median] < x {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(xs) || xs[low] != x {
		return false
	}

	return true
}

func bruteforceSearch(x int, xs []int) bool {
	for _, xss := range xs {
		if xss == x {
			return true
		}
	}

	return false
}

func search(x int, xs []int, fn func(int, []int) bool) bool {
	return fn(x, xs)
}

func main() {
	xs := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	x := 63
	fmt.Println(search(x, xs, binarySearch))
	fmt.Println(search(x, xs, bruteforceSearch))
}
