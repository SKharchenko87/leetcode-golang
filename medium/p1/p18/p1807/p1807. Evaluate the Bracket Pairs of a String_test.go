package p1807

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{s: "(name)is(age)yearsold", knowledge: [][]string{{"name", "bob"}, {"age", "two"}}}, "bobistwoyearsold"},
		{"Example 2", args{s: "hi(name)", knowledge: [][]string{{"a", "b"}}}, "hi?"},
		{"Example 3", args{s: "(a)(a)(a)aaa", knowledge: [][]string{{"a", "yes"}}}, "yesyesyesaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
