package p1106

import "testing"

func Test_parseBoolExpr(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{expression: "&(|(f))"}, false},
		{"Example 2", args{expression: "|(f,f,f,t)"}, true},
		{"Example 3", args{expression: "!(&(f,t))"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
