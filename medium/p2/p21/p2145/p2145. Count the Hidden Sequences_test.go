package p2145

import "testing"

func Test_numberOfArrays(t *testing.T) {
	type args struct {
		differences []int
		lower       int
		upper       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{differences: []int{1, -3, 4}, lower: 1, upper: 6}, 2},
		{"Example 2", args{differences: []int{3, -4, 5, 1, -2}, lower: -4, upper: 5}, 4},
		{"Example 3", args{differences: []int{4, -7, 2}, lower: 3, upper: 6}, 0},
		{"TestCase 64", args{differences: []int{96759, -21770, 16529, -94344, -60558, -19699, 22321, 6815, 63132, -90917}, lower: -84725, upper: 99570}, 4454},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArrays(tt.args.differences, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
