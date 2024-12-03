package d2p1

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay2Part1Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day2Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "2",
			fields: fields{
				daySolverDelegate: &Day2Part1Solver{},
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
				t.Errorf("Day2Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay2Part1Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day2Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "282",
			fields: fields{
				daySolverDelegate: &Day2Part1Solver{},
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
				t.Errorf("Day2Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day2Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReportSafe(t *testing.T) {
	type args struct {
		line      []int
		direction int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Increasing should pass",
			args: args{
				line:      []int{1, 2, 5, 8, 10, 13},
				direction: 1,
			},
			want: true,
		},
		{
			name: "Decreasing should pass",
			args: args{
				line:      []int{13, 10, 8, 5, 2, 1},
				direction: -1,
			},
			want: true,
		},
		{
			name: "Increasing difference greater than 3 should fail",
			args: args{
				line:      []int{1, 2, 5, 8, 10, 14},
				direction: 1,
			},
			want: false,
		},
		{
			name: "Decreasing difference greater than 3 should fail",
			args: args{
				line:      []int{14, 10, 8, 5, 2, 1},
				direction: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReportSafe(tt.args.line, tt.args.direction); got != tt.want {
				t.Errorf("isReportSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
