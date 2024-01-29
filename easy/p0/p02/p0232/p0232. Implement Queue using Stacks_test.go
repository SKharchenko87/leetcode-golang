package p0232

import "testing"

func Test_run(t *testing.T) {
	type args struct {
		command []string
		param   [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.args.command, tt.args.param)
		})
	}
}
