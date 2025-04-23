package p2338

import "testing"

func Test_idealArrays(t *testing.T) {
	type args struct {
		n        int
		maxValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 2, maxValue: 5}, 10},
		{"Example 2", args{n: 5, maxValue: 3}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idealArrays(tt.args.n, tt.args.maxValue); got != tt.want {
				t.Errorf("idealArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
