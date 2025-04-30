package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run cmd/create_day.go <day_number>")
		return
	}
	day := args[1]
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	dayDir := filepath.Join(wd, "day"+day)

	createDirectory(dayDir)

	for i := 1; i <= 2; i++ {
		createDirectory(filepath.Join(dayDir, fmt.Sprintf("p%d", i)))
		createSolverFile(dayDir, day, strconv.Itoa(i))
		createSolverTestFile(dayDir, day, strconv.Itoa(i))
		createInputFile(dayDir, strconv.Itoa(i))
		createInputExampleFile(dayDir, strconv.Itoa(i))
	}
}

func createDirectory(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.Mkdir(dirPath, 0777)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
	}
	fmt.Println("Created directory:", dirPath)
	return nil
}

func createFile(filePath string) (*os.File, error) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil, err
	}
	fmt.Println("Created file:", filePath)
	return file, nil
}

func createSolverFile(dayDir, day, part string) error {
	filePath := filepath.Join(dayDir, "p"+part, "day_"+day+"_part_"+part+"_solver.go")
	file, err := createFile(filePath)
	if err != nil {
		fmt.Println("Error creating solver file:", err)
		return err
	}
	defer file.Close()

	file.WriteString("package d" + day + "p" + part + "\n\n" +
		"type Day" + day + "Part" + part + "Solver struct{}\n\n" +
		"func (solver *Day" + day + "Part" + part + "Solver) Solve(lines []string) (string, error) {\n\n" +
		"\treturn \"\", nil\n" +
		"}\n")

	return nil
}

func createInputFile(dayDir, part string) error {
	filePath := filepath.Join(dayDir, "p"+part, "input.txt")
	file, err := createFile(filePath)
	if err != nil {
		fmt.Println("Error creating input file:", err)
		return err
	}
	defer file.Close()
	return nil
}

func createInputExampleFile(dayDir, part string) error {
	filePath := filepath.Join(dayDir, "p"+part, "input-example.txt")
	file, err := createFile(filePath)
	if err != nil {
		fmt.Println("Error creating input example file:", err)
		return err
	}
	defer file.Close()
	return nil
}

func createSolverTestFile(dayDir, day, part string) error {
	filePath := filepath.Join(dayDir, "p"+part, "day_"+day+"_part_"+part+"_solver_test.go")
	file, err := createFile(filePath)
	if err != nil {
		fmt.Println("Error creating solver test file:", err)
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`package d%sp%s

import (
    "advent-of-code-2024/core"
    "testing"
)

func TestDay%sPart%sSolver_Solve_InputExample(t *testing.T) {
    type fields struct {
        daySolverDelegate *Day%sPart%sSolver
    }
    tests := []struct {
        name    string
        fields  fields
        want    string
        wantErr bool
    }{
        {
            name: "should return correct answer for input example",
            want: "PLACEHOLDER",
            fields: fields{
                daySolverDelegate: &Day%sPart%sSolver{},
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
                t.Errorf("Day%sPart%sSolver.Solve() error = %%v, wantErr %%v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Day%sPart%sSolver.Solve() = %%v, want %%v", got, tt.want)
            }
        })
    }
}

func TestDay%sPart%sSolver_Solve_Input(t *testing.T) {
    type fields struct {
        daySolverDelegate *Day%sPart%sSolver
    }
    tests := []struct {
        name    string
        fields  fields
        want    string
        wantErr bool
    }{
        {
            name: "should return correct answer for input",
            want: "PLACEHOLDER",
            fields: fields{
                daySolverDelegate: &Day%sPart%sSolver{},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
			t.Skip()
            solver := &core.DaySolver{
                DaySolverDelegate: tt.fields.daySolverDelegate,
            }
            got, err := solver.CalculateAnswerFromInput()
            if (err != nil) != tt.wantErr {
                t.Errorf("Day%sPart%sSolver.Solve() error = %%v, wantErr %%v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Day%sPart%sSolver.Solve() = %%v, want %%v", got, tt.want)
            }
        })
    }
}
`, day, part, day, part, day, part, day, part, day, part, day, part, day, part, day, part, day, part, day, part, day, part)
	file.WriteString(content)
	return nil
}
