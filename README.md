# Advent of Code 2024

This repository contains solutions for the Advent of Code 2024 challenges. The project is structured to handle each day's challenge with a dedicated solver that implements the `Solver` interface. The `DaySolver` struct is responsible for reading the input file and calculating the answer using its `Solver` delegate.

## Table of Contents
- [Repository Layout](#repository-layout)
- [Directory and File Descriptions](#directory-and-file-descriptions)
- [Creating Files for a New Day](#creating-files-for-a-new-day)

### Repository Layout
```
cmd/
    create_day.go
core/
    day_solver.go
dayN/
    p1/
        day_N_part_1_solver.go
        day_N_part_1_solver_test.go
        input.txt
    p2/
        day_N_part_2_solver.go
        day_N_part_2_solver_test.go
        input.txt
    day_N_utils.go
go.mod
runChallengeTests.sh
utils/
    file_reader.go
```

### Directory and File Descriptions

- **cmd/**: Contains the [`cmd/create_day.go`](cmd/create_day.go) script used to generate the necessary files for a new day's challenge.
- **core/**: Contains the [`core/day_solver.go`](core/day_solver.go) file which defines the `Solver` interface and the `DaySolver` struct.
- **dayN/**: Contains the solutions and tests for Day N's challenges.
  - **p1/**: Contains the solver, test, and input file for Part 1 of Day N.
  - **p2/**: Contains the solver, test, and input file for Part 2 of Day N.
  - **day_N_utils.go**: Contains utility functions relevant to both parts for a day
- **go.mod**: The Go module file.
- **runChallengeTests.sh**: Shell script to run unit tests for all days
- **utils/**: Contains utility functions, including [`utils/file_reader.go`](utils/file_reader.go) for reading input files.

### `DaySolver` and `Solver` Interface

The `DaySolver` struct is responsible for reading the input file and calculating the answer using its `Solver` delegate. Each day's solver implements the `Solver` interface, which requires a `Solve` method.

### Creating Files for a New Day

The `cmd/create_day.go` script can be used to create all the necessary files for a new day's challenge. Run the script with the day number as an argument:

`go run cmd/create_day.go <day_number>`

This will generate the following files:
```
day<day_number>/
    p1/
        day_<day_number>_part_1_solver.go
        day_<day_number>_part_1_solver_test.go
        input.txt
    p2/
        day_<day_number>_part_2_solver.go
        day_<day_number>_part_2_solver_test.go
        input.txt
```
Each solver file will contain a struct that implements the Solver interface, and the test files will contain basic test cases for the solver.