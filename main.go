package main

import (
    "log"
    "os"
    "strconv"
    "github.com/reshane/aocgo24/days"
)

var solutions = []days.Solver{ days.Day1, days.Day2, days.Day3, days.Day4, days.Day5, days.Day6, days.Day7, days.Day8 }

func main() {
    args := os.Args
    if len(args) < 2 {
        log.Fatal("Error: No day specified")
    }

    day, err := strconv.Atoi(args[1])
    if err != nil {
        log.Fatalf("Error: Invalid day [%s]: %s", args[1], err)
    }

    if day-1 >= len(solutions) || day-1 < 0 {
        log.Fatalf("No implementation for %d", day)
    }
    solutions[day-1].Solve()
}
