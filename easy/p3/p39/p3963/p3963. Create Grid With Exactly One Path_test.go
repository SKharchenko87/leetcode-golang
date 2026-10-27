package p3963

import (
	"reflect"
	"testing"
)

func Test_createGrid(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{m: 2, n: 3}, []string{"...", "##."}},
		{"Example 2", args{m: 3, n: 3}, []string{"...", "##.", "##."}},
		{"Example 3", args{m: 1, n: 4}, []string{"...."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createGrid(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
