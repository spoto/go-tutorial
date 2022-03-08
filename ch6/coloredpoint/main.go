package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	p := ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	p1 := ColoredPoint2{&Point{1, 1}, red}
	q1 := ColoredPoint2{&Point{5, 4}, blue}
	fmt.Println(p1.Distance(*q1.Point))
	q1.Point = p1.Point
	p1.ScaleBy(2)
	fmt.Println(*p1.Point, *q1.Point)

	p2 := Point{1, 2}
	q2 := Point{4, 6}
	distanceFromP := p2.Distance
	fmt.Println(distanceFromP(q2))
	var origin Point
	fmt.Println(distanceFromP(origin))
	scaleP := p2.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p2)

	p3 := Point{1, 2}
	q3 := Point{4, 6}
	distance := Point.Distance
	fmt.Println(distance(p3, q3))
	fmt.Printf("%T\n", distance)
	scale := (*Point).ScaleBy
	scale(&p3, 2)
	fmt.Println(p3)
	fmt.Printf("%T\n", scale)
}
