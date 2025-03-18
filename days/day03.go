package days

import (
    "log"
    "strings"
    "strconv"

    "github.com/reshane/aocgo24/utils"
)

type day3 struct{}
var Day3 day3

func (day3) Part1(contents string) int {
    var ans int = 0
    chunks := strings.Split(contents, "mul(")
    for _, chunk := range chunks {
        // for each chunk
        // we need [numeric] [comma] [numeric] [close paren]
        comma := strings.Index(chunk, ",")
        cParen := strings.Index(chunk, ")")
        if comma != -1 && cParen != -1 {
            left, err := strconv.Atoi(chunk[0:comma])
            if err != nil {
                continue
            }
            right, err := strconv.Atoi(chunk[comma+1:cParen])
            if err != nil {
                continue
            }
            ans += left * right
        }
    }
    return ans
}

func (day3) Part2(contents string) int {
    var ans int = 0
    chunks := strings.Split(contents, "don't()")
    for i, chunk := range chunks {
        if i == 0 {
            ans += Day3.Part1(chunk)
        } else {
            do := strings.Index(chunk, "do()")
            if do != -1 {
                ans += Day3.Part1(chunk[do:])
            }
        }
    }
    return ans
}

func (day3) Solve() {
    log.Print("Day 3 solution")
    contents, err := utils.Io.ReadToString("./input/in3")
    if err != nil {
        return
    }

    log.Printf("Part 1: %d", Day3.Part1(contents))
    log.Printf("Part 2: %d", Day3.Part2(contents))
}

