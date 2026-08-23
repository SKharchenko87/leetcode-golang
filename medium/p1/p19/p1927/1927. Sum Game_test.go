package p1927

import "testing"

func Test_sumGame(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1:", args{num: "5023"}, false},
		{"Example 2:", args{num: "25??"}, true},
		{"Example 3:", args{num: "?3295???"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumGame(tt.args.num); got != tt.want {
				t.Errorf("sumGame() = %v, want %v", got, tt.want)
			}
		})
	}

}
