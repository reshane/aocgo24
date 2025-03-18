package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day7Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t7")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 3749
    if got := Day7.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day7Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t7")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 11387
    if got := Day7.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
