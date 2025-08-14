package p2359

import "testing"

func Test_closestMeetingNode(t *testing.T) {
	type args struct {
		edges []int
		node1 int
		node2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{edges: []int{2, 2, 3, -1}, node1: 0, node2: 1}, 2},
		{"Example 2", args{edges: []int{1, 2, -1}, node1: 0, node2: 2}, 2},
		{"TestCase 74", args{edges: []int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, node1: 5, node2: 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestMeetingNode(tt.args.edges, tt.args.node1, tt.args.node2); got != tt.want {
				t.Errorf("closestMeetingNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 0  1  2   3  4  5  6  7  8  9
// 4, 4, 8, -1, 9, 8, 4, 4, 1, 1

func Benchmark_closestMeetingNode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		closestMeetingNode([]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6)
	}
}

func Benchmark_closestMeetingNode1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		closestMeetingNode1([]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6)
	}
}

func Benchmark_closestMeetingNode0(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		closestMeetingNode0([]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6)
	}
}
