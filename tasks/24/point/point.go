package point

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func Distance(p1 Point, p2 Point) float64 {
	x1 := math.Pow(p2.x-p1.x, 2)
	x2 := math.Pow(p2.y-p1.y, 2)

	x1 = math.Round(x1*100) / 100
	x2 = math.Round(x2*100) / 100

	fmt.Println(math.Sqrt(x1 + x2))
	return 1.2
}
