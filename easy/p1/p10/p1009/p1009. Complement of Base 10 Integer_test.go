package p1009

import "testing"

func Test_bitwiseComplement(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num: 5}, 2},
		{"Example 2", args{num: 7}, 0},
		{"Example 3", args{num: 10}, 5},
		{"TestCase 127", args{num: 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplement(tt.args.num); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
