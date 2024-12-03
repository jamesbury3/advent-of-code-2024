package d1p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestD1P2_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day1Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "31",
			fields: fields{
				daySolverDelegate: &Day1Part2Solver{},
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
				t.Errorf("Day1Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day1Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestD1P2_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day1Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "23150395",
			fields: fields{
				daySolverDelegate: &Day1Part2Solver{},
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
				t.Errorf("Day1Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day1Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
