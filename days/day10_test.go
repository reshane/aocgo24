package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day10Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t10")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 36
    if got := Day10.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day10Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t10")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 81
    if got := Day10.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
