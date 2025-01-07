package p1769

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		boxes string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{boxes: "110"}, []int{1, 1, 3}},
		{"Example 2", args{boxes: "001011"}, []int{11, 8, 5, 4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
