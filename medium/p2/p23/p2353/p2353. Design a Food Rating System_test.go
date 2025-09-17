package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		param    []any
	}
	foodRatingParams := constructorType{[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}, []int{9, 12, 8, 15, 14, 7}}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{"Example 1", args{[]string{"FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"},
			[]any{foodRatingParams, "korean", "japanese", changeRatingType{"sushi", 16}, "japanese", changeRatingType{"ramen", 16}, "japanese"}},
			[]any{nil, "kimchi", "ramen", nil, "sushi", nil, "ramen"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
