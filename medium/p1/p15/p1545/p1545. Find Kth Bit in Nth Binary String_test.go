package p1545

import (
	"reflect"
	"testing"
)

func Test_invert(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "011", args: args{[]byte("011")}, want: []byte("100")},
		{name: "110", args: args{[]byte("110")}, want: []byte("001")},
		{name: "0111001101100011011100100110001", args: args{[]byte("0111001101100011011100100110001")}, want: []byte("1000110010011100100011011001110")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invert(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "011", args: args{[]byte("011")}, want: []byte("110")},
		{name: "110", args: args{[]byte("110")}, want: []byte("011")},
		{name: "1000110010011100100011011001110", args: args{[]byte("1000110010011100100011011001110")}, want: []byte("0111001101100010011100100110001")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthBit(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"Example 1", args{n: 3, k: 1}, '0'},
		{"Example 2", args{n: 4, k: 11}, '1'},
		{"TestCase 7", args{n: 4, k: 12}, '0'},
		{"TestCase 11", args{n: 5, k: 24}, '0'},
		{"My 1", args{n: 5, k: 11}, '1'},
		{"My 2", args{n: 10, k: 11}, '1'},
		{"My 3", args{n: 20, k: 11}, '1'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthBit(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
