package p3751

import "testing"

func Test_totalWaviness(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num1: 120, num2: 130}, 3},
		{"Example 2", args{num1: 198, num2: 202}, 3},
		{"Example 3", args{num1: 4848, num2: 4848}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalWaviness(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("totalWaviness() = %v, want %v", got, tt.want)
			}
		})
	}
}
