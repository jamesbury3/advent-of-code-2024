package d9p2

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay9Part2Solver_Solve_InputExample(t *testing.T) {
    type fields struct {
        daySolverDelegate *Day9Part2Solver
    }
    tests := []struct {
        name    string
        fields  fields
        want    string
        wantErr bool
    }{
        {
            name: "should return correct answer for input example",
            want: "2858",
            fields: fields{
                daySolverDelegate: &Day9Part2Solver{},
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
                t.Errorf("Day9Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Day9Part2Solver.Solve() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDay9Part2Solver_Solve_Input(t *testing.T) {
    type fields struct {
        daySolverDelegate *Day9Part2Solver
    }
    tests := []struct {
        name    string
        fields  fields
        want    string
        wantErr bool
    }{
        {
            name: "should return correct answer for input",
            want: "6547228115826",
            fields: fields{
                daySolverDelegate: &Day9Part2Solver{},
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
                t.Errorf("Day9Part2Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Day9Part2Solver.Solve() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_findStartingIndex(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		space  []string
		length int
		want   int
	}{
		{
			space: []string{".",".",".","."},
			length: 3,
			want: 0,
		},
		{
			space: []string{"1",".",".","."},
			length: 3,
			want: 1,
		},
		{
			space: []string{".",".",".","1"},
			length: 3,
			want: 0,
		},
		{
			space: []string{"4","5",".","1"},
			length: 3,
			want: -1,
		},
		{
			space: []string{"4","5",".",".",".","1"},
			length: 3,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findStartingIndex(tt.space, tt.length)
			if got != tt.want {
				t.Errorf("findStartingIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

