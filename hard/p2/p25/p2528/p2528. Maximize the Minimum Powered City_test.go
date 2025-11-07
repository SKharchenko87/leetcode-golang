package p2528

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		stations []int
		r        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{stations: []int{1, 2, 4, 5, 0}, r: 1, k: 2}, 5},
		{"Example 2", args{stations: []int{4, 4, 4, 4}, r: 0, k: 3}, 4},
		{"TestCase 12", args{stations: []int{4, 2}, r: 1, k: 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.stations, tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
