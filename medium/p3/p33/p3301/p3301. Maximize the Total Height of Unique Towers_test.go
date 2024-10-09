package p3301

import "testing"

func Test_maximumTotalSum(t *testing.T) {
	type args struct {
		maximumHeight []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{maximumHeight: []int{2, 3, 4, 3}}, 10},
		{"Example 2", args{maximumHeight: []int{15, 10}}, 25},
		{"Example 3", args{maximumHeight: []int{2, 2, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalSum(tt.args.maximumHeight); got != tt.want {
				t.Errorf("maximumTotalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
