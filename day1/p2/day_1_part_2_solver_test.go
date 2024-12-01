package d1p2

import (
	"advent-of-code-2025/utils"
	"testing"
)

func TestD1P2_Solve(t *testing.T) {
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
			name: "Solve Problem",
			want: "23150395",
			fields: fields{
				daySolverDelegate: &Day1Part2Solver{},
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
				t.Errorf("Day1Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day1Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
