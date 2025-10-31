package p3289

import (
	"reflect"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"Example 1", args{nums: []int{0, 1, 1, 0}}, []int{0, 1}},
		{"Example 2", args{nums: []int{0, 3, 2, 1, 3, 2}}, []int{2, 3}},
		{"Example 3", args{nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}}, []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := getSneakyNumbers(tt.args.nums); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("getSneakyNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
