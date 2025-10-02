package p2523

import (
	"reflect"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{left: 10, right: 19}, []int{11, 13}},
		{"Example 2", args{left: 4, right: 6}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestPrimes(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{n: 1}, true},
		{"2", args{n: 2}, true},
		{"3", args{n: 3}, true},
		{"4", args{n: 4}, false},
		{"5", args{n: 5}, true},
		{"6", args{n: 6}, false},
		{"7", args{n: 7}, true},
		{"8", args{n: 8}, false},
		{"9", args{n: 9}, false},
		{"10", args{n: 10}, false},
		{"11", args{n: 11}, true},
		{"12", args{n: 12}, false},
		{"13", args{n: 13}, true},
		{"14", args{n: 14}, false},
		{"15", args{n: 15}, false},
		{"16", args{n: 16}, false},
		{"17", args{n: 17}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
