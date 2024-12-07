package d7p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay7Part2Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day7Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "11387",
			fields: fields{
				daySolverDelegate: &Day7Part2Solver{},
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
				t.Errorf("Day7Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day7Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay7Part2Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day7Part2Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "340362529351427",
			fields: fields{
				daySolverDelegate: &Day7Part2Solver{},
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
				t.Errorf("Day7Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day7Part2Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combineFloat64s(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "should return 12 for 1 and 2",
			args: args{
				a: 1,
				b: 2,
			},
			want: 12,
		},
		{
			name: "should return 5678 for 5 and 678",
			args: args{
				a: 5,
				b: 678,
			},
			want: 5678,
		},
		{
			name: "should return 1234567 for 123 and 4567",
			args: args{
				a: 123,
				b: 4567,
			},
			want: 1234567,
		},
		{
			name: "should return 12345670 for 1234567 and 0",
			args: args{
				a: 1234567,
				b: 0,
			},
			want: 12345670,
		},
		{
			name: "should return 1234567 for 0 and 1234567",
			args: args{
				a: 0,
				b: 1234567,
			},
			want: 1234567,
		},
		{
			name: "should return 1234567 for 12345 and 67",
			args: args{
				a: 12345,
				b: 67,
			},
			want: 1234567,
		},
		{
			name: "should return 0 for 0 and 0",
			args: args{
				a: 0,
				b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineFloat64s(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("combineInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
