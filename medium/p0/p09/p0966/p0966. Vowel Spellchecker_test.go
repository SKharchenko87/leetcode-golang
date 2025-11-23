package p0966

import (
	"reflect"
	"testing"
)

func Test_spellchecker(t *testing.T) {
	type args struct {
		wordlist []string
		queries  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{wordlist: []string{"KiTe", "kite", "hare", "Hare"}, queries: []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}}, []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}},
		{"Example 2", args{wordlist: []string{"yellow"}, queries: []string{"YellOw"}}, []string{"yellow"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spellchecker(tt.args.wordlist, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
