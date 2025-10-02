package p3304

import "testing"

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{k: 5}, 'b'},
		{"Example 2", args{k: 10}, 'c'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthCharacter(tt.args.k); got != tt.want {
				t.Errorf("kthCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
