package p0066

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{digits: []int{1, 2, 3}}, []int{1, 2, 4}},
		{"Example 2", args{digits: []int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"Example 3", args{digits: []int{9}}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_plusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOne([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3})
	}
}

func Benchmark_plusOne0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOne0([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 1, 2, 3})
	}
}
