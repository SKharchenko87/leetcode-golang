package p2452

import (
	"reflect"
	"testing"
)

func Test_twoEditWords(t *testing.T) {
	type args struct {
		queries    []string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{queries: []string{"word", "note", "ants", "wood"}, dictionary: []string{"wood", "joke", "moat"}}, []string{"word", "note", "wood"}},
		{"Example 2", args{queries: []string{"yes"}, dictionary: []string{"not"}}, []string{}},
		{"My 1", args{queries: []string{"aoxqockwtrbodtkefuwusiqvgyvrkfluiexyvvbuylsbgkdmpsas"}, dictionary: []string{"aoxqockwtrbodtkefuwusiqvgyvrkfluiexyvvbuylsbgkdmpsfs"}}, []string{"aoxqockwtrbodtkefuwusiqvgyvrkfluiexyvvbuylsbgkdmpsas"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEditWords(tt.args.queries, tt.args.dictionary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoEditWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
