package p2043

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		params   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes []interface{}
	}{
		{"Example 1", args{[]string{"Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"}, []interface{}{[]int64{10, 100, 20, 50, 30}, []int64{3, 10}, []int64{5, 1, 20}, []int64{5, 20}, []int64{3, 4, 15}, []int64{10, 50}}}, []interface{}{nil, true, true, true, false, false}},
		{"Example 1", args{[]string{"Bank", "deposit", "transfer", "transfer"}, []interface{}{[]int64{0}, []int64{1, 2}, []int64{1, 1, 1}, []int64{1, 1, 3}}}, []interface{}{nil, true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := run(tt.args.commands, tt.args.params); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("run() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
