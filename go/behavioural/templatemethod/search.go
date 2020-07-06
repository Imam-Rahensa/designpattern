package main

import "fmt"

type isearch interface {
	doSearch(x int, xs []int) bool
}

type search struct {
	isearch isearch
}

func (s *search) caridong(x int, xs []int) bool {
	fmt.Println("Searching..")
	return s.isearch.doSearch(x, xs)
}
