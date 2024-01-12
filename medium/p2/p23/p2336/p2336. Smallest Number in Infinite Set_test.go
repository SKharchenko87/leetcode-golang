package p2336

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		commands []string
		args     [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1",
			args{
				[]string{"SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"},
				[][]int{{}, {2}, {}, {}, {}, {1}, {}, {}, {}},
			},
			[]int{null, null, 1, 2, 3, null, 1, 4, 5},
		},
		{"Case 96",
			args{
				[]string{"SmallestInfiniteSet", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest"},
				[][]int{{}, {}, {1}, {}, {}, {}, {2}, {3}, {}, {}},
			},
			[]int{null, 1, null, 1, 2, 3, null, null, 2, 3},
		},
		{"Case 101",
			args{
				[]string{"SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "addBack", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "popSmallest", "popSmallest", "addBack", "addBack", "popSmallest", "addBack", "addBack", "addBack", "addBack", "addBack"},
				[][]int{{}, {942}, {}, {}, {769}, {}, {286}, {}, {943}, {}, {}, {263}, {}, {}, {}, {}, {156}, {}, {}, {873}, {141}, {}, {962}, {}, {}, {}, {}, {447}, {972}, {26}, {233}, {844}, {396}, {}, {772}, {}, {}, {431}, {578}, {}, {204}, {}, {135}, {}, {709}, {}, {}, {}, {}, {}, {}, {418}, {}, {976}, {250}, {}, {}, {}, {891}, {66}, {}, {751}, {}, {743}, {}, {}, {}, {}, {}, {363}, {715}, {}, {750}, {451}, {460}, {635}, {83}, {}, {}, {400}, {}, {}, {}, {}, {497}, {}, {18}, {}, {719}, {}, {}, {243}, {}, {}, {}, {}, {224}, {}, {}, {}, {}, {}, {984}, {}, {}, {91}, {468}, {}, {36}, {}, {}, {}, {}, {302}, {108}, {}, {809}, {441}, {495}, {335}, {}, {}, {847}, {773}, {}, {}, {}, {981}, {}, {453}, {165}, {}, {339}, {415}, {786}, {434}, {483}, {}, {434}, {948}, {10}, {552}, {519}, {560}, {}, {370}, {509}, {59}, {}, {507}, {}, {974}, {899}, {}, {}, {}, {}, {592}, {779}, {336}, {}, {858}, {}, {}, {}, {}, {826}, {97}, {728}, {}, {427}, {}, {}, {}, {}, {192}, {}, {973}, {806}, {130}, {}, {}, {}, {}, {}, {173}, {}, {558}, {49}, {82}, {620}, {880}, {902}, {969}, {192}, {48}, {}, {}, {385}, {219}, {436}, {}, {650}, {}, {426}, {}, {128}, {}, {432}, {750}, {}, {852}, {112}, {463}, {608}, {706}, {120}, {351}, {}, {889}, {}, {837}, {880}, {}, {55}, {}, {}, {609}, {301}, {}, {607}, {791}, {862}, {176}, {35}},
			},
			[]int{null, null, 1, 2, null, 3, null, 4, null, 5, 6, null, 7, 8, 9, 10, null, 11, 12, null, null, 13, null, 14, 15, 16, 17, null, null, null, null, null, null, 18, null, 19, 20, null, null, 21, null, 22, null, 23, null, 24, 25, 26, 27, 28, 29, null, 30, null, null, 31, 32, 33, null, null, 34, null, 35, null, 36, 37, 38, 39, 40, null, null, 41, null, null, null, null, null, 42, 43, null, 44, 45, 46, 47, null, 48, null, 18, null, 49, 50, null, 51, 52, 53, 54, null, 55, 56, 57, 58, 59, null, 60, 61, null, null, 62, null, 36, 63, 64, 65, null, null, 66, null, null, null, null, 67, 68, null, null, 69, 70, 71, null, 72, null, null, 73, null, null, null, null, null, 74, null, null, null, null, null, null, 10, null, null, null, 59, null, 75, null, null, 76, 77, 78, 79, null, null, null, 80, null, 81, 82, 83, 84, null, null, null, 85, null, 86, 87, 88, 89, null, 90, null, null, null, 91, 92, 93, 94, 95, null, 96, null, null, null, null, null, null, null, null, null, 48, 49, null, null, null, 82, null, 97, null, 98, null, 99, null, null, 100, null, null, null, null, null, null, null, 101, null, 102, null, null, 103, null, 55, 104, null, null, 105, null, null, null, null, null},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.commands, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
