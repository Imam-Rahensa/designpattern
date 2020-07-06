package main

type bruteforce struct {
	search search
}

func (b *bruteforce) doSearch(x int, xs []int) bool {
	for _, xss := range xs {
		if xss == x {
			return true
		}
	}

	return false
}
