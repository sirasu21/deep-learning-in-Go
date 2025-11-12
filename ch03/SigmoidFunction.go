package ch03

import "math"


func SigmoidFunction(x float64) float64{
	return 1 / (1 + math.Exp(-x))
}
