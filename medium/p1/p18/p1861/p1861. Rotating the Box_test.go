package p1861

import (
	"reflect"
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	type args struct {
		box [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"Example 1", args{box: [][]byte{{'#', '.', '#'}}}, [][]byte{{'.'}, {'#'}, {'#'}}},
		{"Example 2", args{box: [][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}}, [][]byte{{'#', '.'}, {'#', '#'}, {'*', '*'}, {'.', '.'}}},
		{"Example 3", args{box: [][]byte{{'#', '#', '*', '.', '*', '.'}, {'#', '#', '#', '*', '.', '.'}, {'#', '#', '#', '.', '#', '.'}}}, [][]byte{{'.', '#', '#'}, {'.', '#', '#'}, {'#', '#', '*'}, {'#', '*', '.'}, {'#', '.', '*'}, {'#', '.', '.'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateTheBox(tt.args.box); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
