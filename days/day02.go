package days

import (
    "log"
    "strings"
    "strconv"
    "slices"

    "github.com/reshane/aocgo24/utils"
)

type day2 struct{}
var Day2 day2

func (day2) ParseInput(lines []string) ([][]int, error) {
    reports := make([][]int, 0)
    for _, line := range lines {
        nums := strings.Split(line, " ")

        levels := make([]int, 0)
        for _, num := range nums {
            level, err := strconv.Atoi(num)
            if err != nil {
                log.Printf("Could not convert %s from [%s] to int: %s", nums[0], line, err)
                return reports, err
            }
            levels = append(levels, level)
        }
        reports = append(reports, levels)
    }

    return reports, nil
}

func (day2) Part1(reports [][]int) int {
    var ans int = 0
    for _, levels := range reports {
        if checkLevels(levels) {
            ans += 1
        }
    }

    return ans
}

func checkLevels(levels []int) bool {
    diffs := make([]int, 0)
    safe := true
    for i := 1; i < len(levels); i++ {
        // append a 1 for positive diff
        // append a 0 for negative diff
        diff := levels[i-1] - levels[i]
        if diff < 0 {
            diffs = append(diffs, 0)
        } else {
            diffs = append(diffs, 1)
        }
        if diff > 3 || diff < -3 || diff == 0 {
            // bypass later check
            safe = false
            break
        }
    }
    if safe {
        for i := 1; i < len(diffs); i++ {
            if diffs[i-1] != diffs[i] {
                safe = false
                break
            }
        }
    }
    return safe
}

func (day2) Part2(reports [][]int) int {
    var ans int = 0
    // for every bad report, if remove each level and try again
    for _, levels := range reports {
        if checkLevels(levels) {
            ans += 1
        } else {
            for i := 0; i < len(levels); i++ {
                iRemoved := slices.Concat(levels[0:i], levels[i+1:])
                if checkLevels(iRemoved) {
                    ans += 1
                    break
                }
            }
        }
    }
    return ans
}

func (day2) Solve() {
    log.Printf("Day 2 solution\n")
    data, err := utils.Io.ReadToLines("./input/in2")
    if err != nil {
        return
    }

    reports, err := Day2.ParseInput(data)
    if err != nil {
        log.Fatalf("Could not parse input: %s", err)
    }

    log.Printf("Part 1: %d", Day2.Part1(reports))
    log.Printf("Part 2: %d", Day2.Part2(reports))
}
