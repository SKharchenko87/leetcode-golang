package p0726

import "testing"

func Test_countOfAtoms(t *testing.T) {
	type args struct {
		formula string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{formula: "H2O"}, "H2O"},
		{"Example 2", args{formula: "Mg(OH)2"}, "H2MgO2"},
		{"Example 3", args{formula: "K4(ON(SO3)2)2"}, "K4N2O14S4"},
		{"My 1", args{formula: "Mg(OH10)2"}, "H20MgO2"},
		{"My 2", args{formula: "K4(O(X10F)2N(SO3)2)2"}, "F4K4N2O14S4X40"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOfAtoms(tt.args.formula); got != tt.want {
				t.Errorf("countOfAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}
