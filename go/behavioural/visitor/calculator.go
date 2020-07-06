package main

import (
	"fmt"
	"math"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square", s.side*s.side)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	//Calculate are for circle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for circle", math.Phi*float32(s.radius)*float32(s.radius))
}

func (a *areaCalculator) visitForrectangle(s *rectangle) {
	//Calculate are for rectangle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for rectangle", s.b*s.l)
}
