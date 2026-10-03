package p1096

import (
	"reflect"
	"testing"
)

func Test_braceExpansionII(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{expression: "{a,b}{c,{d,e}}"}, []string{"ac", "ad", "ae", "bc", "bd", "be"}},
		{"Example 2", args{expression: "{{a,z},a{b,c},{ab,z}}"}, []string{"a", "ab", "ac", "z"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := braceExpansionII(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
