package p2418

import (
	"reflect"
	"testing"
)

func Test_sortPeople(t *testing.T) {
	type args struct {
		names   []string
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{names: []string{"Mary", "John", "Emma"}, heights: []int{180, 165, 170}}, []string{"Mary", "Emma", "John"}},
		{"Example 2", args{names: []string{"Alice", "Bob", "Bob"}, heights: []int{155, 185, 150}}, []string{"Bob", "Alice", "Bob"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeople(tt.args.names, tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
