package p0399

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"Case 1", args{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
		{"Case 2", args{[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}}, []float64{3.75000, 0.40000, 5.00000, 0.20000}},
		{"Case 3", args{[][]string{{"a", "b"}}, []float64{0.5}, [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}}}, []float64{0.50000, 2.00000, -1.00000, -1.00000}},
		{"Case 4", args{[][]string{{"a", "e"}, {"b", "e"}}, []float64{4.0, 3.0}, [][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}}}, []float64{1.3333333333333333, 1.00000, -1.00000}},
		{"Case 4", args{[][]string{{"a", "b"}, {"c", "d"}}, []float64{1.0, 1.0}, [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}}}, []float64{-1.00000, -1.00000, 1.00000, 1.00000}},
		{"Case 16", args{[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}}, []float64{360.00000, 0.008333333333333333, 20.00000, 1.00000, -1.00000, -1.00000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
