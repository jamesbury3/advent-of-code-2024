package d5p1

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay5Part1Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day5Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "143",
			fields: fields{
				daySolverDelegate: &Day5Part1Solver{},
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
				t.Errorf("Day5Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay5Part1Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day5Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "4996",
			fields: fields{
				daySolverDelegate: &Day5Part1Solver{},
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
				t.Errorf("Day5Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
