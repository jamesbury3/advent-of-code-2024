package d2p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay2Part2Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day2Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "4",
			fields: fields{
				daySolverDelegate: &Day2Part2Solver{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &core.DaySolver{
				DaySolverDelegate: tt.fields.daySolverDelegate,
			}
			got, err := solver.CalculateAnswerFromInputExample()
			if (err != nil) != tt.wantErr {
				t.Errorf("Day2Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay2Part2Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day2Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "349",
			fields: fields{
				daySolverDelegate: &Day2Part2Solver{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &core.DaySolver{
				DaySolverDelegate: tt.fields.daySolverDelegate,
			}
			got, err := solver.CalculateAnswerFromInput()
			if (err != nil) != tt.wantErr {
				t.Errorf("Day2Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReportSafeWithOneRemoval(t *testing.T) {
	type args struct {
		line               []int
		idx                int
		direction          int
		prev               *int
		removalHasOccurred bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Increasing no removals should pass",
			args: args{
				line:               []int{1, 2, 5, 8, 10, 13},
				idx:                1,
				direction:          1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: true,
		},
		{
			name: "Decreasing no removals should pass",
			args: args{
				line:               []int{13, 10, 8, 5, 2, 1},
				idx:                1,
				direction:          -1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: true,
		},
		{
			name: "Increasing one removal should pass",
			args: args{
				line:               []int{1, 2, 5, -15, 8, 10},
				idx:                1,
				direction:          1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: true,
		},
		{
			name: "Decreasing one removal should pass",
			args: args{
				line:               []int{13, 10, 29, 8, 5, 2},
				idx:                1,
				direction:          -1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: true,
		},
		{
			name: "Increasing two removals should fail",
			args: args{
				line:               []int{1, 2, 5, -15, 7, 11},
				idx:                1,
				direction:          1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: false,
		},
		{
			name: "Decreasing two removals should fail",
			args: args{
				line:               []int{13, 10, 29, 8, -5, 6},
				idx:                1,
				direction:          -1,
				prev:               nil,
				removalHasOccurred: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReportSafeWithOneRemoval(tt.args.line, tt.args.idx, tt.args.direction, tt.args.prev, tt.args.removalHasOccurred); got != tt.want {
				t.Errorf("isReportSafeWithOneRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReportSafeInAnyOrder(t *testing.T) {
	type args struct {
		line []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Increasing with removal of first should pass",
			args: args{
				line: []int{59, 2, 5, 8, 10, 13},
			},
			want: true,
		},
		{
			name: "Decreasing with removal of first should pass",
			args: args{
				line: []int{13, 10, 8, 5, 2, 1},
			},
			want: true,
		},
		{
			name: "Increasing with removal of last should pass",
			args: args{
				line: []int{1, 2, 5, 8, 10, 9},
			},
			want: true,
		},
		{
			name: "Decreasing with removal of last should pass",
			args: args{
				line: []int{20, 18, 15, 13, 11, 15},
			},
			want: true,
		},
		{
			name: "Increasing with removal of second to last should pass",
			args: args{
				line: []int{1, 2, 5, 8, -10, 11},
			},
			want: true,
		},
		{
			name: "Decreasing with removal of second to last should pass",
			args: args{
				line: []int{13, 10, 8, 5, 49, 2},
			},
			want: true,
		},
		{
			name: "Increasing with removal of second should pass",
			args: args{
				line: []int{1, -50, 3, 5, 8, 11},
			},
			want: true,
		},
		{
			name: "Decreasing with removal of second should pass",
			args: args{
				line: []int{13, 50, 10, 8, 5, 3},
			},
			want: true,
		},
		{
			name: "Increasing with 2 unsafe reports should fail",
			args: args{
				line: []int{59, 2, 5, -4, 10, 14},
			},
			want: false,
		},
		{
			name: "Decreasing with 2 unsafe reports should fail",
			args: args{
				line: []int{14, 10, 8, 5, -2, 1},
			},
			want: false,
		},
		{
			name: "Increasing with first two the same should pass",
			args: args{
				line: []int{5, 5, 6, 8, 10, 13},
			},
			want: true,
		},
		{
			name: "Decreasing with first two the same should pass",
			args: args{
				line: []int{13, 10, 8, 6, 5, 5},
			},
			want: true,
		},
		{
			name: "Increasing with last two the same should pass",
			args: args{
				line: []int{1, 2, 5, 8, 10, 10},
			},
			want: true,
		},
		{
			name: "Decreasing with last two the same should pass",
			args: args{
				line: []int{13, 10, 8, 5, 2, 2},
			},
			want: true,
		},
		{
			name: "Decreasing with removal of second and duplicate values should pass",
			args: args{
				line: []int{28, 25, 26, 25, 24},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReportSafeInAnyOrder(tt.args.line); got != tt.want {
				t.Errorf("isReportSafeInAnyOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
