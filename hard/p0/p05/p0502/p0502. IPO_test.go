package p0502

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}

	profit17, err := os.ReadFile("profit17.data")
	if err != nil {
		log.Fatal(err)
	}
	profit17ArrStr := strings.Split(strings.Trim(string(profit17), "[]"), ",")
	profit17Arr := make([]int, len(profit17ArrStr))
	for i, s := range profit17ArrStr {
		profit17Arr[i], _ = strconv.Atoi(s)
	}

	capital17, err := os.ReadFile("capital17.data")
	if err != nil {
		log.Fatal(err)
	}
	capital17ArrStr := strings.Split(strings.Trim(string(capital17), "[]"), ",")
	capital17Arr := make([]int, len(capital17ArrStr))
	for i, s := range capital17ArrStr {
		capital17Arr[i], _ = strconv.Atoi(s)
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{k: 2, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 1, 1}}, 4},
		{"Example 2", args{k: 3, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 1, 2}}, 6},
		{"Example 3", args{k: 1, w: 0, profits: []int{1, 2, 3}, capital: []int{1, 1, 2}}, 0},
		{"Example 8", args{k: 2, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 9, 10}}, 1},
		{"Example 17", args{k: 111, w: 12, profits: profit17Arr, capital: capital17Arr}, 120730},
		{"Example 19", args{k: 1, w: 2, profits: []int{1, 2, 3}, capital: []int{1, 1, 2}}, 5},
		{"My 1", args{k: 3, w: 0, profits: []int{1, 2, 3}, capital: []int{0, 1, 1}}, 6},
		{"My 1", args{k: 3, w: 0, profits: []int{3, 2, 1}, capital: []int{1, 1, 0}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
