package ch02

func AddGate(x1, x2 float64) int {
	w1, w2, theta := 0.5, 0.5, 0.7
	tmp := x1*w1 + x2*w2
	if tmp <= theta {
		return 0
	}
	return 1
}
