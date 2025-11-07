package ch02

func XorGate(x1, x2 float64) int{
	s1 := NandGate(x1, x2)
	s2 := OrGate(x1, x2)
	y := AndGate(float64(s1), float64(s2))
	return y
}