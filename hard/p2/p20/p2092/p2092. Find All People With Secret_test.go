package p2092

import (
	"reflect"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
	type args struct {
		n           int
		meetings    [][]int
		firstPerson int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{n: 6, meetings: [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, firstPerson: 1}, []int{0, 1, 2, 3, 5}},
		{"Example 2", args{n: 4, meetings: [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, firstPerson: 3}, []int{0, 1, 3}},
		{"Example 3", args{n: 5, meetings: [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, firstPerson: 1}, []int{0, 1, 2, 3, 4}},
		{"TestCase 10", args{n: 6, meetings: [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}}, firstPerson: 1}, []int{0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllPeople(tt.args.n, tt.args.meetings, tt.args.firstPerson); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
