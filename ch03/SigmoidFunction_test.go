package ch03

import "testing"

func TestSigmoidFunction(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{0}, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SigmoidFunction(tt.args.x); got != tt.want {
				t.Errorf("SigmoidFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
