package p1371

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{s: "eleetminicoworoep"}, 13},
		{"Example 1", args{s: "eleetminicoworoepoie"}, 18},
		{"Example 1", args{s: "xxxeleetminicoworoepoie"}, 21},
		{"Example 1", args{s: "xxxeleetminicoworoep"}, 13},
		{"Example 1", args{s: "xxxxleetminicoworoep"}, 17},
		{"My 1", args{s: "eleetemn"}, 8},
		{"Example 2", args{s: "leetcodeisgreat"}, 5},
		{"Example 3", args{s: "bcbcbc"}, 6},
		{"TestCase 36", args{s: "id"}, 1},
		{"TestCase 27", args{s: "yopumzgd"}, 4},
		{"My 2", args{s: "aei"}, 0},
		{"My 3", args{s: "axxxee"}, 5},
		{"My 3", args{s: "eaxxxe"}, 3},
		{"My 3", args{s: "eaxxxexx"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

//eleetmini

//itminicoworoepoie
//eleetminicoworoepo

//xxxeleetminicoworoepoie
//xxxeleetminicoworoepo

//xxxeleetminicoworoep
//    leetminicowor

//xxxxleetminicoworoep
//xxxxleetminicowor

//eleetminicoworoepoie
//eleetminicoworoepo
