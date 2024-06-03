package p0344

import (
	"reflect"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Example 1", args{[]byte{'h', 'e', 'l', 'l', 'o'}}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"Example 2", args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
