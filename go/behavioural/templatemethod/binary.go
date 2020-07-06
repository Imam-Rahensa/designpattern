package main

type binary struct {
	search search
}

func (b *binary) doSearch(x int, xs []int) bool {
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
