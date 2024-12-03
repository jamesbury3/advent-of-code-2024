package d3p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay3Part2Solver_Solve(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day3Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Solve Problem",
			want: "83158140",
			fields: fields{
				daySolverDelegate: &Day3Part2Solver{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &core.DaySolver{
				DaySolverDelegate: tt.fields.daySolverDelegate,
			}
			got, err := solver.CalculateAnswer()
			if (err != nil) != tt.wantErr {
				t.Errorf("Day3Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day3Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
