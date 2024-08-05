package p2053

import "testing"

func Test_kthDistinct(t *testing.T) {
	type args struct {
		arr []string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{arr: []string{"d", "b", "c", "b", "c", "a"}, k: 2}, "a"},
		{"Example 2", args{arr: []string{"aaa", "aa", "a"}, k: 1}, "aaa"},
		{"Example 3", args{arr: []string{"a", "b", "a"}, k: 3}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthDistinct(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kthDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
