package p0670

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{num: 2736}, 7236},
		{"Example 2", args{num: 9973}, 9973},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setNumber(t *testing.T) {
	type args struct {
		num int
		i   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{num: 2736, i: 0, x: 0}, 2730},
		{"1", args{num: 2736, i: 1, x: 0}, 2706},
		{"2", args{num: 2736, i: 2, x: 0}, 2036},
		{"3", args{num: 2736, i: 3, x: 0}, 736},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setNumber(tt.args.num, tt.args.i, tt.args.x); got != tt.want {
				t.Errorf("setNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
