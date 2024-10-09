package p0641

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		cmd    []string
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"Example 1", args{
			[]string{"MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"},
			[][]int{{3}, {1}, {2}, {3}, {4}, {}, {}, {}, {4}, {}},
		}, []interface{}{nil, true, true, true, false, 2, true, true, true, 4}},
		{"Example 1", args{
			[]string{"MyCircularDeque", "insertFront", "deleteLast", "getRear", "getFront", "getFront", "deleteFront", "insertFront", "insertLast", "insertFront", "getFront", "insertFront"},
			[][]int{{4}, {9}, {}, {}, {}, {}, {}, {6}, {5}, {9}, {}, {6}},
		}, []interface{}{nil, true, true, -1, -1, -1, false, true, true, true, 9, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.cmd, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularDeque_IsFull(t *testing.T) {
	type fields struct {
		queue []int
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"1", fields{make([]int, 5), 0, 3}, true},
		{"2", fields{make([]int, 5), 1, 4}, true},
		{"3", fields{make([]int, 5), 2, 0}, true},
		{"4", fields{make([]int, 5), 3, 1}, true},
		{"5", fields{make([]int, 5), 4, 2}, true},
		{"6", fields{make([]int, 5), 5, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MyCircularDeque{
				queue: tt.fields.queue,
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := this.IsFull(); got != tt.want {
				t.Errorf("IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
