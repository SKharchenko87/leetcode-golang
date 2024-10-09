package p1813

import "testing"

func Test_areSentencesSimilar(t *testing.T) {
	type args struct {
		sentence1 string
		sentence2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{sentence1: "My name is Haley", sentence2: "My Haley"}, true},
		{"Example 2", args{sentence1: "of", sentence2: "A lot of words"}, false},
		{"Example 3", args{sentence1: "Eating right now", sentence2: "Eating"}, true},
		{"TestCase 97", args{sentence1: "Luky", sentence2: "Lucccky"}, false},
		{"TestCase 100", args{sentence1: "Are You Okay", sentence2: "are you okay"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.args.sentence1, tt.args.sentence2); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
