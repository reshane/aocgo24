package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day12Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t12")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 140
    if got := Day12.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day12Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t12.2")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 368
    if got := Day12.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
