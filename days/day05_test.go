package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day5Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t5")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 143
    if got := Day5.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day5Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t5")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 123
    if got := Day5.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
