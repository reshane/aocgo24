package days

import (
    "log"
    "testing"
    "github.com/reshane/aocgo24/utils"
)

func Test_day2Part1(t *testing.T) {
    lines, err := utils.Io.ReadToLines("../input/t2")
    if err != nil {
        log.Fatal(err)
    }
    reports, err := Day2.ParseInput(lines)
    if err != nil {
        log.Fatal(err)
    }
    expected := 2
    if got := Day2.Part1(reports); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day2Part2(t *testing.T) {
    lines, err := utils.Io.ReadToLines("../input/t2")
    if err != nil {
        log.Fatal(err)
    }
    reports, err := Day2.ParseInput(lines)
    if err != nil {
        log.Fatal(err)
    }
    expected := 4
    if got := Day2.Part2(reports); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
