package main

import "fmt"

func main() {
	xs := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	x := 63

	binary := &binary{}
	s := search{
		isearch: binary,
	}
	fmt.Println(s.caridong(x, xs))

	bruteforce := &bruteforce{}
	s = search{
		isearch: bruteforce,
	}
	fmt.Println(s.caridong(x, xs))
}
