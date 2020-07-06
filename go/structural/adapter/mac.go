package main

import "fmt"

type mac struct {
}

func (m *mac) insertUSB() {
	fmt.Println("Insert square port into mac machine")
}
