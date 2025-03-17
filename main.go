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

    if day-1 >= len(days.Solutions) || day-1 < 0 {
        log.Fatalf("No implementation for %d", day)
    }
    days.Solutions[day-1].Solve()
}
