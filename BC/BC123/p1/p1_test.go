package p1

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{[]int{3, 4, 5}}, "scalene"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
