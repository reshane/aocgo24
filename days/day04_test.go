package days

import (
    "testing"

    "github.com/reshane/aocgo24/utils"
)

func Test_day4Part1(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t4")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 18
    if got := Day4.Part1(contents); got != expected {
        t.Errorf("Part 1 = %v, expected %v", got, expected)
    }
}

func Test_day4Part2(t *testing.T) {
    contents, err := utils.Io.ReadToString("../input/t4")
    if err != nil {
        t.Error("Could not get test input")
    }

    expected := 9
    if got := Day4.Part2(contents); got != expected {
        t.Errorf("Part 2 = %v, expected %v", got, expected)
    }
}
