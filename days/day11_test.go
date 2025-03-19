package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day11Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t11")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 55312
    if got := Day11.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day11Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t11")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 65601038650482
    if got := Day11.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
