package p2751

import (
	"reflect"
	"testing"
)

func Test_survivedRobotsHealths(t *testing.T) {
	type args struct {
		positions  []int
		healths    []int
		directions string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{positions: []int{5, 4, 3, 2, 1}, healths: []int{2, 17, 9, 15, 10}, directions: "RRRRR"}, []int{2, 17, 9, 15, 10}},
		{"Example 2", args{positions: []int{3, 5, 2, 6}, healths: []int{10, 10, 15, 12}, directions: "RLRL"}, []int{14}},
		{"Example 3", args{positions: []int{1, 2, 5, 6}, healths: []int{10, 10, 11, 11}, directions: "RLRL"}, []int{}},
		{"Testcase 42", args{positions: []int{1, 40}, healths: []int{10, 11}, directions: "RL"}, []int{10}},
		{"Testcase 348", args{positions: []int{11, 44, 16}, healths: []int{1, 20, 17}, directions: "RLR"}, []int{18}},
		{"Testcase 560", args{positions: []int{34, 50, 42, 2}, healths: []int{6, 27, 17, 38}, directions: "LLRR"}, []int{36}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := survivedRobotsHealths(tt.args.positions, tt.args.healths, tt.args.directions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("survivedRobotsHealths() = %v, want %v", got, tt.want)
			}
		})
	}
}
