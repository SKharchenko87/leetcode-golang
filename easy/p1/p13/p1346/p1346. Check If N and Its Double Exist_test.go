package p1346

import "testing"

func Test_checkIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{arr: []int{10, 2, 5, 3}}, true},
		{"Example 2", args{arr: []int{3, 1, 7, 11}}, false},
		{"My 1", args{arr: []int{3, 1, 7, 0}}, false},
		{"My 2", args{arr: []int{0, 1, 7, 0}}, true},
		{"Test case 2", args{arr: []int{-2, -2, 0, 0, 10, -19, 4, 6, -8}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
