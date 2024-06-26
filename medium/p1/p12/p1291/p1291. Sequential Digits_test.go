package p1291

import (
	"reflect"
	"testing"
)

func Test_sequentialDigits(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{100, 300}, []int{123, 234}},
		{"Case 2", args{1000, 13000}, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
		{"Case 2", args{58, 155}, []int{67, 78, 89, 123}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequentialDigits(tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sequentialDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
