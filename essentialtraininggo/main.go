package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrNegSqrt    = errors.New("sqrt of negative number")
	ErrNoSolution = errors.New("no solution found")
)

func Abs(val float64) float64 {
	if val < 0 {
		return -val
	}
	return val
}

func Sqrt(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, fmt.Errorf("error")
	}
	return math.Sqrt(val), nil
}
func main() {

}
