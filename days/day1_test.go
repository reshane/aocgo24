package days

import (
    "log"
    "testing"
    "github.com/reshane/aocgo24/utils"
)

func Test_day1Part1(t *testing.T) {
    lines, err := utils.Io.ReadToLines("../input/t1")
    if err != nil {
        log.Fatal(err)
    }
    left, right, err := Day1.ParseInput(lines)
    if err != nil {
        log.Fatal(err)
    }
    expected := 11
    if got := Day1.Part1(left, right); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day1Part2(t *testing.T) {
    lines, err := utils.Io.ReadToLines("../input/t1")
    if err != nil {
        log.Fatal(err)
    }
    left, right, err := Day1.ParseInput(lines)
    if err != nil {
        log.Fatal(err)
    }
    expected := 31
    if got := Day1.Part2(left, right); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}
