package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day6Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t6")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 41
    if got := Day6.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day6Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t6")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 6
    if got := Day6.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
