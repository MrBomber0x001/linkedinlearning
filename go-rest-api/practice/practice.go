package practice

import "fmt"

type Point struct {
	X      int
	Y      int
	Length int
}

func (p *Point) CalculateLength() int {
	return p.X ^ 2 - p.Y ^ 2
}

func Test() {
	p := Point{X: 1, Y: 2}
	l := p.CalculateLength()
	fmt.Println(l)
}
