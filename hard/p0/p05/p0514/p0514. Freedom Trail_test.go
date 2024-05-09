package p0514

import "testing"

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"godding", "gd"}, 4},
		{"Example 2", args{"godding", "godding"}, 13},
		{"Example a0", args{"xrrakuulnczywjs", "jrlucwzakzuss"}, 58},
		{"Example a0", args{"xrrakuulnczywjs", "jrlucwzakzussrlckyjjsuwkua"}, 119},
		{"Example 46", args{"xrrakuulnczywjs", "jrlucwzakzussrlckyjjsuwkuarnaluxnyzcnrxxwruyr"}, 204},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotateSteps(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
