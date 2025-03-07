package main

import (
    "log"
    "os"
    "strconv"
    "github.com/reshane/aocgo24/days"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        log.Fatal("Error: No day specified")
    }

    day, err := strconv.Atoi(args[1])
    if err != nil {
        log.Fatalf("Error: Invalid day [%s]: %s", args[1], err)
    }
    solutions := []func(){days.Day1.Solve, days.Day2.Solve, days.Day3.Solve}
    if day-1 >= len(solutions) {
        log.Fatalf("No implementation for %d", day)
    }
    solutions[day-1]()
}
