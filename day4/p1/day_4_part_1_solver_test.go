package d4p1

import (
	"advent-of-code-2024/core"
	"testing"
)

func TestDay4Part1Solver_Solve_InputExample(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day4Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input example",
			want: "18",
			fields: fields{
				daySolverDelegate: &Day4Part1Solver{},
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
				t.Errorf("Day4Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day4Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay4Part1Solver_Solve_Input(t *testing.T) {
	type fields struct {
		daySolverDelegate *Day4Part1Solver
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "should return correct answer for input",
			want: "2549",
			fields: fields{
				daySolverDelegate: &Day4Part1Solver{},
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
				t.Errorf("Day4Part1Solver.Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Day4Part1Solver.Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLeft(t *testing.T) {
	type args struct {
		row     int
		col     int
		letters [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false if col is less than 3",
			args: args{
				row:     0,
				col:     2,
				letters: [][]rune{},
			},
			want: false,
		},
		{
			name: "should return true if letters match XMAS",
			args: args{
				row:     0,
				col:     3,
				letters: [][]rune{{'X', 'M', 'A', 'S'}},
			},
			want: true,
		},
		{
			name: "should return true if letters match SAMX",
			args: args{
				row:     0,
				col:     3,
				letters: [][]rune{{'S', 'A', 'M', 'X'}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLeft(tt.args.row, tt.args.col, tt.args.letters); got != tt.want {
				t.Errorf("checkLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUp(t *testing.T) {
	type args struct {
		row     int
		col     int
		letters [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false if row is less than 3",
			args: args{
				row:     2,
				col:     0,
				letters: [][]rune{},
			},
			want: false,
		},
		{
			name: "should return true if letters match XMAS",
			args: args{
				row:     3,
				col:     0,
				letters: [][]rune{{'X'}, {'M'}, {'A'}, {'S'}},
			},
			want: true,
		},
		{
			name: "should return true if letters match SAMX",
			args: args{
				row:     3,
				col:     0,
				letters: [][]rune{{'S'}, {'A'}, {'M'}, {'X'}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUp(tt.args.row, tt.args.col, tt.args.letters); got != tt.want {
				t.Errorf("checkUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkLeftUpDiagonal(t *testing.T) {
	type args struct {
		row     int
		col     int
		letters [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false if row is less than 3",
			args: args{
				row:     2,
				col:     0,
				letters: [][]rune{},
			},
			want: false,
		},
		{
			name: "should return false if col is less than 3",
			args: args{
				row:     0,
				col:     2,
				letters: [][]rune{},
			},
			want: false,
		},
		{
			name: "should return true if letters match XMAS",
			args: args{
				row:     3,
				col:     3,
				letters: [][]rune{{'X'}, {'.', 'M'}, {'.', '.', 'A'}, {'.', '.', '.', 'S'}},
			},
			want: true,
		},
		{
			name: "should return true if letters match SAMX",
			args: args{
				row:     3,
				col:     3,
				letters: [][]rune{{'S'}, {'.', 'A'}, {'.', '.', 'M'}, {'.', '.', '.', 'X'}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLeftUpDiagonal(tt.args.row, tt.args.col, tt.args.letters); got != tt.want {
				t.Errorf("checkLeftUpDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkUpRightDiagonal(t *testing.T) {
	type args struct {
		row     int
		col     int
		letters [][]rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return false if row is less than 3",
			args: args{
				row:     2,
				col:     0,
				letters: [][]rune{},
			},
			want: false,
		},
		{
			name: "should return false if col is greater than or equal to len(letters[0])-3",
			args: args{
				row:     0,
				col:     0,
				letters: [][]rune{{'X'}},
			},
			want: false,
		},
		{
			name: "should return true if letters match XMAS",
			args: args{
				row:     3,
				col:     0,
				letters: [][]rune{{'.', '.', '.', 'X'}, {'.', '.', 'M', ','}, {'.', 'A', '.', '.'}, {'S', '.', '.', '.'}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkUpRightDiagonal(tt.args.row, tt.args.col, tt.args.letters); got != tt.want {
				t.Errorf("checkUpRightDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
