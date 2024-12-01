package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	createDirectory(filepath.Join(dayDir, "p1"))
	createDirectory(filepath.Join(dayDir, "p2"))

	createSolverFile(dayDir, day, "1")
	createSolverFile(dayDir, day, "2")

	createSolverTestFile(dayDir, day, "1")
	createSolverTestFile(dayDir, day, "2")

	createInputFile(dayDir, "1")
	createInputFile(dayDir, "2")
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
    "advent-of-code-2025/utils"
    "testing"
)

func TestDay%sPart%sSolver_Solve(t *testing.T) {
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
            name: "Solve Problem",
            want: "",
            fields: fields{
                daySolverDelegate: &Day%sPart%sSolver{},
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
                t.Errorf("Day%sPart%sSolver.Solve() error = %%v, wantErr %%v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Day%sPart%sSolver.Solve() = %%v, want %%v", got, tt.want)
            }
        })
    }
}
`, day, part, day, part, day, part, day, part, day, part, day, part)
	file.WriteString(content)
	return nil
}
