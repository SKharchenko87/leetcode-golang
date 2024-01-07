package p1137

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{4}, 4},
		{"Case 5", args{5}, 7},
		{"Case 6", args{6}, 13},
		{"Case 7", args{7}, 24},
		{"Case 8", args{8}, 44},
		{"Case 9", args{9}, 81},
		{"Case 10", args{10}, 149},
		{"Case 11", args{11}, 274},
		{"Case 12", args{12}, 504},
		{"Case 13", args{13}, 927},
		{"Case 14", args{14}, 1705},
		{"Case 15", args{15}, 3136},
		{"Case 16", args{16}, 5768},
		{"Case 17", args{17}, 10609},
		{"Case 18", args{18}, 19513},
		{"Case 19", args{19}, 35890},
		{"Case 20", args{20}, 66012},
		{"Case 21", args{21}, 121415},
		{"Case 22", args{22}, 223317},
		{"Case 23", args{23}, 410744},
		{"Case 24", args{24}, 755476},
		{"Case 25", args{25}, 1389537},
		{"Case 26", args{26}, 2555757},
		{"Case 27", args{27}, 4700770},
		{"Case 28", args{28}, 8646064},
		{"Case 29", args{29}, 15902591},
		{"Case 30", args{30}, 29249425},
		{"Case 31", args{31}, 53798080},
		{"Case 32", args{32}, 98950096},
		{"Case 33", args{33}, 181997601},
		{"Case 34", args{34}, 334745777},
		{"Case 35", args{35}, 615693474},
		{"Case 36", args{36}, 1132436852},
		{"Case 37", args{37}, 2082876103},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
