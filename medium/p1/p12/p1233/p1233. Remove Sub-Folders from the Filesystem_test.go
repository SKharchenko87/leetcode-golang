package p1233

import (
	"reflect"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	type args struct {
		folder []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}}, []string{"/a", "/c/d", "/c/f"}},
		{"Example 2", args{folder: []string{"/a", "/a/b/c", "/a/b/d"}}, []string{"/a"}},
		{"Example 3", args{folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}}, []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeSubfolders(tt.args.folder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.want)
			}
		})
	}
}
