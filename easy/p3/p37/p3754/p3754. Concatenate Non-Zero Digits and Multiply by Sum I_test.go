package p3754

import "testing"

func Test_sumAndMultiply(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Example 1", args{n: 10203004}, 12340},
		{"Example 2", args{n: 1000}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAndMultiply(tt.args.n); got != tt.want {
				t.Errorf("sumAndMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
		wantSum int
	}{
		{"12034", args{12034}, 1234, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotSum := solution(tt.args.n)
			if gotRes != tt.wantRes {
				t.Errorf("solution() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotSum != tt.wantSum {
				t.Errorf("solution() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
