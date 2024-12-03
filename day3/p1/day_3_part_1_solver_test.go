package d3p1

import (
	"advent-of-code-2024/core"
	"strconv"
	"testing"
)

func TestDay3Part1Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day3Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "161",
			fields: fields{
				daySolverDelegate: &Day3Part1Solver{},
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
				t.Errorf("Day3Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day3Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3Part1Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day3Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "173785482",
			fields: fields{
				daySolverDelegate: &Day3Part1Solver{},
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
				t.Errorf("Day3Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day3Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay3Part1Solver_Solve2(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day3Part1Solver
	}
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "muls on multiple lines should return sum of products",
			args: args{
				lines: []string{
					"m()mul(198,449)#from(",
					"()$mul(381,350)",
				},
			},
			fields: fields{
				daySolverDelegate: &Day3Part1Solver{},
			},
			want: strconv.Itoa((198 * 449) + (381 * 350)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solver := &core.DaySolver{
				DaySolverDelegate: tt.fields.daySolverDelegate,
			}
			got, err := solver.DaySolverDelegate.Solve(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("Day3Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day3Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProductFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "no mnul should return 0",
			args: args{
				line: "ul(1,2)",
			},
			want: 0,
		},
		{
			name: "one mul should return product",
			args: args{
				line: "m()mul(198,449)#from(",
			},
			want: 198 * 449,
		},
		{
			name: "two muls should return sum of products",
			args: args{
				line: "()$mul(381,350)m()mul(198,449)#from(",
			},
			want: (381 * 350) + (198 * 449),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfProductsFromLine(tt.args.line); got != tt.want {
				t.Errorf("getProductFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
