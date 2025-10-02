package p1980

import "testing"

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{nums: []string{"01", "10"}}, "11"},
		{"Example 2", args{nums: []string{"00", "01"}}, "11"},
		{"Example 3", args{nums: []string{"111", "011", "001"}}, "101"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifferentBinaryString(tt.args.nums); got != tt.want {
				t.Errorf("findDifferentBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
