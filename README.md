# advent-of-code-2024

This repository contains my solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges.

I am going to solve the challenges using Go. This is my second edition of the Advent of Code.

## Requirements

- Go 1.21 or higher
- pre-commit
- commitizen (optional)

## Running the solutions

To run the solutions, you need to have Go installed on your machine. You can download it [here](https://golang.org/dl/).

To run a solution, run the following command from the root of the project:

```bash
go run dayXX/main.go
```

Replace `XX` with the day number you want to run.

This will output something like this:

```bash
Result XXa: XXX
Result XXb: XXX
```

Where `XX` is the day number and `XXX` is the result of the solution.

## Project architecture

Inputs are stored in the `inputs` directory. Each day has its own directory, named `dayXX`, where `XX` is the day number. Each day directory contains a `main.go` file with the solution for that day.
Inputs should not be pushed to the repository, so they are ignored in the `.gitignore` file.

```bash
.
├── LICENSE
├── README.md
├── day01
│   └── main.go
├── go.mod
└── inputs
    ├── input01a.txt
    └── input01b.txt
```

## Total time spent

- Total time spent: 5 hours 40 minutes
