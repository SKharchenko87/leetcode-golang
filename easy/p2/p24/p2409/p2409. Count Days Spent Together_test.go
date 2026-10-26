package p2409

import "testing"

func Test_countDaysTogether(t *testing.T) {
	type args struct {
		arriveAlice string
		leaveAlice  string
		arriveBob   string
		leaveBob    string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{arriveAlice: "08-15", leaveAlice: "08-18", arriveBob: "08-16", leaveBob: "08-19"}, 3},
		{"Example 2", args{arriveAlice: "10-01", leaveAlice: "10-31", arriveBob: "11-01", leaveBob: "12-31"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDaysTogether(tt.args.arriveAlice, tt.args.leaveAlice, tt.args.arriveBob, tt.args.leaveBob); got != tt.want {
				t.Errorf("countDaysTogether() = %v, want %v", got, tt.want)
			}
		})
	}
}
