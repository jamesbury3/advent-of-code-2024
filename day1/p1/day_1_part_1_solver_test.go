package d1p1

import (
	"advent-of-code-2025/utils"
	"testing"
)

func TestDay1Part1Solver_Solve(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day1Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Solve Problem",
			want: "936063",
			fields: fields{
				daySolverDelegate: &Day1Part1Solver{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &utils.DaySolver{
				DaySolverDelegate: tt.fields.daySolverDelegate,
			}
			got, err := solver.CalculateAnswer()
			if (err != nil) != tt.wantErr {
				t.Errorf("Day1Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day1Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
