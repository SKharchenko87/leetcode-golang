package p1462

import (
	"reflect"
	"testing"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"Example 1", args{numCourses: 2, prerequisites: [][]int{{1, 0}}, queries: [][]int{{0, 1}, {1, 0}}}, []bool{false, true}},
		{"Example 2", args{numCourses: 2, prerequisites: [][]int{}, queries: [][]int{{1, 0}, {0, 1}}}, []bool{false, false}},
		{"Example 3", args{numCourses: 3, prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}}, queries: [][]int{{1, 0}, {1, 2}}}, []bool{true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfPrerequisite(tt.args.numCourses, tt.args.prerequisites, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkIfPrerequisite() = %v, want %v", got, tt.want)
			}
		})
	}
}
