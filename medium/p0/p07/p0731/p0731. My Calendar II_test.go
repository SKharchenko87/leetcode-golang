package p0731

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
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
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}},
			[]bool{true, true, true, false, true, true},
		},
		{"TestCase 1", args{
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {26, 35}, {26, 32}, {25, 32}, {18, 26}, {40, 45}, {19, 26}, {48, 50}, {1, 6}, {46, 50}, {11, 18}}},
			[]bool{true, true, false, true, true, true, true, true, true, true},
		},
		{"TestCase 73", args{
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {33, 44}, {85, 95}, {20, 37}, {91, 100}, {89, 100}, {77, 87}, {80, 95}, {42, 61}, {40, 50}, {85, 99}, {74, 91}, {70, 82}, {5, 17}, {77, 89}, {16, 26}, {21, 31}, {30, 43}, {96, 100}, {27, 39}, {44, 55}, {15, 34}, {85, 99}, {74, 93}, {84, 94}, {82, 94}, {46, 65}, {31, 49}, {58, 73}, {86, 99}, {73, 84}, {68, 80}, {5, 18}, {75, 87}, {88, 100}, {25, 41}, {66, 79}, {28, 41}, {60, 70}, {62, 73}, {16, 33}}},
			[]bool{true, true, true, true, false, true, false, true, false, false, false, true, true, false, true, false, false, true, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{"My 1", args{
			[]string{"MyCalendarTwo", "book", "book", "book"},
			[][2]int{{}, {4, 8}, {1, 4}, {4, 5}}},
			[]bool{true, true, true},
		},
		{"My 2", args{
			[]string{"MyCalendarTwo", "book", "book", "book"},
			[][2]int{{}, {4, 8}, {1, 4}, {1, 5}}},
			[]bool{true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.cmd, tt.args.books); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}},
			[]bool{true, true, true, false, true, true},
		},
		{"TestCase 1", args{
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {26, 35}, {26, 32}, {25, 32}, {18, 26}, {40, 45}, {19, 26}, {48, 50}, {1, 6}, {46, 50}, {11, 18}}},
			[]bool{true, true, false, true, true, true, true, true, true, true},
		},
		{"TestCase 73", args{
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {33, 44}, {85, 95}, {20, 37}, {91, 100}, {89, 100}, {77, 87}, {80, 95}, {42, 61}, {40, 50}, {85, 99}, {74, 91}, {70, 82}, {5, 17}, {77, 89}, {16, 26}, {21, 31}, {30, 43}, {96, 100}, {27, 39}, {44, 55}, {15, 34}, {85, 99}, {74, 93}, {84, 94}, {82, 94}, {46, 65}, {31, 49}, {58, 73}, {86, 99}, {73, 84}, {68, 80}, {5, 18}, {75, 87}, {88, 100}, {25, 41}, {66, 79}, {28, 41}, {60, 70}, {62, 73}, {16, 33}}},
			[]bool{true, true, true, true, false, true, false, true, false, false, false, true, true, false, true, false, false, true, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{"TestCase 47", args{
			[]string{"MyCalendarTwo", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book", "book"},
			[][2]int{{}, {97, 100}, {51, 65}, {27, 46}, {90, 100}, {20, 32}, {15, 28}, {60, 73}, {77, 91}, {67, 85}, {58, 72}, {74, 93}, {73, 83}, {71, 87}, {97, 100}}},
			[]bool{true, true, true, true, true, false, true, true, true, false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.cmd, tt.args.books); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
