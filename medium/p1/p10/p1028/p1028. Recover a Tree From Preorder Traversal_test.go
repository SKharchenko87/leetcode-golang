package p1028

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		traversal string
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example 1", args{traversal: "1-2--3--4-5--6--7"}, []any{1, 2, 5, 3, 4, 6, 7}},
		{"Example 2", args{traversal: "1-2--3---4-5--6---7"}, []any{1, 2, 5, 3, nil, 6, nil, 4, nil, 7}},
		{"Example 3", args{traversal: "1-401--349---90--88"}, []any{1, 401, nil, 349, 88, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.traversal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
