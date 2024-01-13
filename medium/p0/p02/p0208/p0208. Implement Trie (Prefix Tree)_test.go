package p0208

import (
	"reflect"
	"testing"
)

func Test_p0208(t *testing.T) {
	type args struct {
		commands []string
		args     [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{name: "Case 1", args: args{[]string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}, [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}},
			want: [][]bool{nil, nil, {true}, {false}, {true}, nil, {true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p0208(tt.args.commands, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("p0208() = %v, want %v", got, tt.want)
			}
		})
	}
}
