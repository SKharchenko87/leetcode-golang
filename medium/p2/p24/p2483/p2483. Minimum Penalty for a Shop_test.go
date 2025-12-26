package p2483

import "testing"

func Test_bestClosingTime(t *testing.T) {
	type args struct {
		customers string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{customers: "YYNY"}, 2},
		{"Example 2", args{customers: "NNNNN"}, 0},
		{"Example 3", args{customers: "YYYY"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestClosingTime(tt.args.customers); got != tt.want {
				t.Errorf("bestClosingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
