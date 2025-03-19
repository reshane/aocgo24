package days

import (
    "log"
    "fmt"
    "strings"
    "strconv"

    "github.com/reshane/aocgo24/utils"
)

type day11 struct{}
var Day11 day11

func (day11) Solve() {
    log.Printf("Day 11 Solutions:\n")
    contents, err := utils.Io.ReadToString("./input/in11")
    if err != nil {
        return
    }
    log.Printf("Part 1: %d\n", Day11.Part1(contents))
    log.Printf("Part 2: %d\n", Day11.Part2(contents))
}

func (day11) parseInput(contents string) map[int]int {
    stones := make(map[int]int, 0)
    for _, stoneString := range strings.Split(contents, " ") {
        stone, err := strconv.Atoi(strings.TrimSpace(stoneString))
        if err != nil {
            log.Fatalf("Could not parse int from string %s\n", stoneString)
        }
        if ct, exists := stones[stone]; exists {
            stones[stone] = ct + 1
        } else {
            stones[stone] = 1
        }
    }
    return stones
}

func (day11) stepStone(stone int) []int {
    if stone == 0 {
        return []int{ 1 }
    }
    stoneString := fmt.Sprintf("%d", stone)
    if len(stoneString) % 2 == 0 {
        left, err := strconv.Atoi(stoneString[0:(len(stoneString)/2)])
        if err != nil {
            log.Fatalf("Could not parse int from string %s\n", stoneString[0:(len(stoneString)/2)])
        }
        right, err := strconv.Atoi(stoneString[(len(stoneString)/2):])
        if err != nil {
            log.Fatalf("Could not parse int from string %s\n", stoneString[(len(stoneString)/2):])
        }
        return []int{left, right}
    }
    return []int{ stone * 2024 }
}

func (day11) stepN(stones map[int]int, iterations int) int {
    freqs := stones
    for i := 0; i < iterations; i++ {
        next := make(map[int]int)
        for k, v := range freqs {
            stepped := Day11.stepStone(k)
            
            for _, step := range stepped {
                if ct, exists := next[step]; exists {
                    next[step] = ct + v
                } else {
                    next[step] = v
                }
            }
        }
        freqs = next
    }

    var ans int = 0
    for _, v := range freqs {
        ans += v
    }
    return ans
}

func (day11) Part1(contents string) int {
    stones := Day11.parseInput(contents)
    return Day11.stepN(stones, 25)
}

func (day11) Part2(contents string) int {
    stones := Day11.parseInput(contents)
    return Day11.stepN(stones, 75)
}
