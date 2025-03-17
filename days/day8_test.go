package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day8Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t8")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 14
    if got := Day8.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day8Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t8")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 34
    if got := Day8.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
