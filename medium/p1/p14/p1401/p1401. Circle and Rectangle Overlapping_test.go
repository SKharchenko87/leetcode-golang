package p1401

import "testing"

func Test_checkOverlap(t *testing.T) {
	type args struct {
		radius  int
		xCenter int
		yCenter int
		x1      int
		y1      int
		x2      int
		y2      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{radius: 1, xCenter: 0, yCenter: 0, x1: 1, y1: -1, x2: 3, y2: 1}, true},
		{"Example 2", args{radius: 1, xCenter: 1, yCenter: 1, x1: 1, y1: -3, x2: 2, y2: -1}, false},
		{"Example 3", args{radius: 1, xCenter: 0, yCenter: 0, x1: -1, y1: 0, x2: 0, y2: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOverlap(tt.args.radius, tt.args.xCenter, tt.args.yCenter, tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("checkOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
