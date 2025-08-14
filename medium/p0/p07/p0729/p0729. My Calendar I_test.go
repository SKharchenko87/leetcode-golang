package p0729

import (
	"reflect"
	"testing"
)

func Test_run1(t *testing.T) {
	type args struct {
		cmd   []string
		books [][2]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"Example 1", args{
			[]string{"MyCalendar", "book", "book", "book"},
			[][2]int{{}, {10, 20}, {15, 25}, {20, 30}},
		}, []bool{true, false, true}},
		{"Example 8", args{
			[]string{"MyCalendar", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {47, 50}, {33, 41}, {39, 45}, {33, 42}, {25, 32}, {26, 35}, {19, 25}, {3, 8}, {8, 13}, {18, 27}},
		}, []bool{true, true, false, false, true, false, true, true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.cmd, tt.args.books); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
