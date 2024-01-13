package p0901

import (
	"reflect"
	"testing"
)

func Test_p0901(t *testing.T) {
	type args struct {
		comands []string
		values  [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{[]string{"StockSpanner", "next", "next", "next", "next", "next", "next", "next"}, [][]int{{}, {100}, {80}, {60}, {70}, {60}, {75}, {85}}}, []int{1, 1, 1, 2, 1, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p0901(tt.args.comands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("p0901() = %v, want %v", got, tt.want)
			}
		})
	}
}
