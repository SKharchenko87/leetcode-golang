package p2115

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findAllRecipes(t *testing.T) {
	type args struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1", args{recipes: []string{"bread"}, ingredients: [][]string{{"yeast", "flour"}}, supplies: []string{"yeast", "flour", "corn"}}, []string{"bread"}},
		{"Example 2", args{recipes: []string{"bread", "sandwich"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}}, supplies: []string{"yeast", "flour", "meat"}}, []string{"bread", "sandwich"}},
		{"Example 3", args{recipes: []string{"bread", "sandwich", "burger"}, ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}, supplies: []string{"yeast", "flour", "meat"}}, []string{"bread", "sandwich", "burger"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllRecipes(tt.args.recipes, tt.args.ingredients, tt.args.supplies)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllRecipes() = %v, want %v", got, tt.want)
			}
		})
	}
}
