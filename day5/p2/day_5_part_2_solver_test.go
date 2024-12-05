package d5p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay5Part2Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day5Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "123",
			fields: fields{
				daySolverDelegate: &Day5Part2Solver{},
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
				t.Errorf("Day5Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay5Part2Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day5Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "6311",
			fields: fields{
				daySolverDelegate: &Day5Part2Solver{},
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
				t.Errorf("Day5Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day5Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
