package p1718

import (
	"reflect"
	"testing"
)

func Test_constructDistancedSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 3}, []int{3, 1, 2, 3, 2}},
		{"Example 2", args{n: 5}, []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructDistancedSequence(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructDistancedSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
