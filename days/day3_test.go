package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day3Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t3.1")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 161
    if got := Day3.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day3Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t3.2")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 48
    if got := Day3.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
