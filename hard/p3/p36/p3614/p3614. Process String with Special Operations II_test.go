package p3614

import "testing"

func Test_processStr(t *testing.T) {
	type args struct {
		s string
		k int64
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{s: "a#b%*", k: 1}, 'a'},
		{"Example 2", args{s: "cd%#*#", k: 3}, 'd'},
		{"Example 3", args{s: "z*#", k: 0}, '.'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("processStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
