package p2109

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_aa(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 string
	}{
		{"Test 1 - Borodino", args{"Borodino.txt"}, []int{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw, _ := os.ReadFile(tt.args.filename)
			data := string(raw)
			arr := getIndexSpaces(data)
			fmt.Println(arr)
			ds := deleteSpace(data)
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("%v", arr)
			}
			if ds != tt.want1 {
				t.Errorf("%v", ds)
			}
		})
	}
}
