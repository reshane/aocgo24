package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day9Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t9")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 1928
    if got := Day9.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day9Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t9")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 2858
    if got := Day9.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
