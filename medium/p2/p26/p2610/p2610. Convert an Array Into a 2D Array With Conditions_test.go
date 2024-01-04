package p2610

import (
	"reflect"
	"testing"
)

func Test_findMatrix(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Case 1", args{[]int{1, 3, 4, 1, 2, 3, 1}}, [][]int{{1, 3, 4, 2}, {1, 3}, {1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMatrix(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
