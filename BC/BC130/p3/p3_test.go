package p3

import "testing"

func Test_abs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.args.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumSubstringsInPartition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "fabccddg"}, 3},
		{"Example 2", args{s: "abababaccddb"}, 2},
		{"TestCase 310", args{s: "aab"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSubstringsInPartition(tt.args.s); got != tt.want {
				t.Errorf("minimumSubstringsInPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
