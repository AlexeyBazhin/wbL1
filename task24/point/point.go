package point

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 //данные поля инкапсулированы - они доступны лишь в рамках данного пакета, т.к. написаны со строчной буквы
}

func NewPoint(x, y float64) *Point {
	point := &Point{}
	point.SetX(x)
	point.SetY(y)
	return point
}

// зачастую, если нужна возможность доступа к инкапсулированым полям извне,
// то используют геттеры-сеттеры, внутри которых может содержаться некая логика
func (point *Point) GetX() float64 {
	return point.x
}
func (point *Point) GetY() float64 {
	return point.y
}
func (point *Point) SetX(x float64) {
	if x < 0 {
		fmt.Println("Отрицательное значение")
	}
	point.x = x
}
func (point *Point) SetY(y float64) {
	if y < 0 {
		fmt.Println("Отрицательное значение")
	}
	point.y = y
}
func (point *Point) DistanceTo(anotherPoint *Point) float64 {
	dx := point.x - anotherPoint.x
	dy := point.y - anotherPoint.y
	return math.Sqrt(dx*dx + dy*dy)
}
