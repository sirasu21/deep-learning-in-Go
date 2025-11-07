package ch02

import "gonum.org/v1/gonum/mat"

func AndGate(x1, x2 float64) int {
	x := mat.NewVecDense(2, []float64{x1, x2})
	w := mat.NewVecDense(2, []float64{0.5, 0.5})
	b := -0.7
	tmp := mat.Dot(x, w) + b
	if tmp <= 0 {
		return 0
	}
	return 1
}
