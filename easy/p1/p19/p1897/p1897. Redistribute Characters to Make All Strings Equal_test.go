package p1897

import "testing"

func Test_makeEqual(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{[]string{"abc", "aabc", "bc"}}, true},
		{"Case 1", args{[]string{"ab", "a"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeEqual(tt.args.words); got != tt.want {
				t.Errorf("makeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
