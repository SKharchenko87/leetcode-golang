package p3606

import (
	"reflect"
	"testing"
)

func Test_validateCoupons(t *testing.T) {
	type args struct {
		code         []string
		businessLine []string
		isActive     []bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{code: []string{"SAVE20", "", "PHARMA5", "SAVE@20"}, businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"}, isActive: []bool{true, true, true, true}}, []string{"PHARMA5", "SAVE20"}},
		{"Example 2", args{code: []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"}, businessLine: []string{"grocery", "electronics", "invalid"}, isActive: []bool{false, true, true}}, []string{"ELECTRONICS_50"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCoupons(tt.args.code, tt.args.businessLine, tt.args.isActive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateCoupons() = %v, want %v", got, tt.want)
			}
		})
	}
}
