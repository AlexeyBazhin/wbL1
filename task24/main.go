package main

import (
	"fmt"
	"wbL1/task24/point"
)

func main() {
	p1 := point.NewPoint(1.0, -2.0)
	p2 := point.NewPoint(4.0, 2.0)
	distance := p1.DistanceTo(p2)
	fmt.Printf("Расстояние между точкой (%.2f, %.2f) и точкой (%.2f, %.2f) равняется %.2f\n", p1.GetX(), p1.GetY(), p2.GetX(), p2.GetY(), distance)
}
