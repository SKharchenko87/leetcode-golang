package p1882

import (
	"reflect"
	"testing"
)

func Test_assignTasks(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		servers []int
		tasks   []int
		want    []int
	}{
		{"Example 1", []int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}, []int{2, 2, 0, 2, 1, 2}},
		{"Example 2", []int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}, []int{1, 4, 1, 4, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := assignTasks(tt.servers, tt.tasks)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
