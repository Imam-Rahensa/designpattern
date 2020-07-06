package main

func main()  {
	areaCalculator := &areaCalculator{}

	circle := &circle{
		radius: 7,
	}
	circle.accept(areaCalculator)
}