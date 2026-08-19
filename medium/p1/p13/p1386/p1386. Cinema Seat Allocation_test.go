package p1386

import "testing"

func Test_maxNumberOfFamilies(t *testing.T) {
	type args struct {
		n             int
		reservedSeats [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{n: 3, reservedSeats: [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}}, 4},
		{"Example 2", args{n: 2, reservedSeats: [][]int{{2, 1}, {1, 8}, {2, 6}}}, 2},
		{"Example 3", args{n: 4, reservedSeats: [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}}, 4},
		{"TestCase 41", args{n: 3, reservedSeats: [][]int{{2, 3}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfFamilies(tt.args.n, tt.args.reservedSeats); got != tt.want {
				t.Errorf("maxNumberOfFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}
