package days

import (
    "log"
    "strings"
    "strconv"
    "sort"

    "github.com/reshane/aocgo24/utils"
)

type day1 struct{}
var Day1 day1

func (day1) ParseInput(lines []string) ([]int, []int, error) {
    left := make([]int, 0, 1000)
    right := make([]int, 0, 1000)
    for _, line := range lines {
        nums := strings.Split(line, "   ")
        
        numLeft, err := strconv.Atoi(nums[0])
        if err != nil {
            log.Printf("Could not convert %s from [%s] to int: %s", nums[0], line, err)
            return left, right, err
        }
        left = append(left, numLeft)
        numRight, err := strconv.Atoi(nums[1])
        if err != nil {
            log.Printf("Could not convert %s from [%s] to int: %s", nums[1], line, err)
            return left, right, err
        }
        right = append(right, numRight)
    }

    sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
    sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

    return left, right, nil
}

func (day1) Part1(left, right []int) int {
    var ans int = 0
    for i := 0; i < len(left); i++ {
        total := left[i] - right[i]
        if total < 0 {
            total *= -1
        }
        ans += total
    }
    return ans
}

func (day1) Part2(left, right []int) int {
    freq := make(map[int]int)
    for _, val := range right {
        count, exists := freq[val]
        if exists {
            freq[val] = count + 1
        } else {
            freq[val] = 1
        }
    }

    var ans int = 0
    for _, val := range left {
        count, exists := freq[val]
        if exists {
            ans += val * count
        }
    }

    return ans
}

func (day1) Solve() {
    log.Printf("Day 1 solution\n")
    data, err := utils.Io.ReadToLines("./input/in1")
    if err != nil {
        return
    }

    left, right, err := Day1.ParseInput(data)
    if err != nil {
        log.Fatalf("Could not parse input: %s", err)
    }

    log.Printf("Part 1: %d", Day1.Part1(left, right))
    log.Printf("Part 2: %d", Day1.Part2(left, right))
}
