package days

import (
    "log"
    "strings"

    "github.com/reshane/aocgo24/utils"
)

type day4 struct{}
var Day4 day4

type point struct {
    x int
    y int
}

func (this *point) add(other point) {
    this.x += other.x
    this.y += other.y
}

func (day4) ParseInput(contents string) ([]point, []point, map[point]string) {
    wordSearch := make(map[point]string)
    xes := make([]point, 0)
    as := make([]point, 0)
    lines := strings.Split(contents, "\n")
    for y, line := range lines {
        for x := 0; x < len(line); x++ {
            if line[x:x+1] == "X" {
                xes = append(xes, point{ x, y })
            }
            if line[x:x+1] == "A" {
                as = append(as, point{ x, y })
            }
            wordSearch[point{ x, y }] = line[x:x+1]
        }
    }
    return xes, as, wordSearch
}

func (day4) Part1(contents string) int {
    var ans int = 0
    dirs := []point{ {-1,-1}, {0,-1}, {1,-1}, {-1,0}, {1,0}, {-1,1}, {0,1}, {1,1} }
    xes, _, wordSearch := Day4.ParseInput(contents)
    for _, start := range xes {
        for _, dir := range dirs {
            curr := start
            curr.add(dir)
            if val, exists := wordSearch[curr]; !exists || val != "M" {
                continue
            }
            curr.add(dir)
            if val, exists := wordSearch[curr]; !exists || val != "A" {
                continue
            }
            curr.add(dir)
            if val, exists := wordSearch[curr]; !exists || val != "S" {
                continue
            }
            ans += 1
        }
    }
    return ans
}

func (this *point) plus(other point) point {
    return point{ this.x + other.x, this.y + other.y }
}

func (day4) Part2(contents string) int {
    var ans int = 0
    _, as, wordSearch := Day4.ParseInput(contents)
    for _, start := range as {
        topRight, trExists := wordSearch[start.plus(point{1,-1})]
        bottomLeft, blExists := wordSearch[start.plus(point{-1,1})]

        topLeft, tlExists := wordSearch[start.plus(point{-1,-1})]
        bottomRight, brExists := wordSearch[start.plus(point{1,1})]

        firstDiag := trExists && blExists && (topRight != bottomLeft) && (topRight == "M" || topRight == "S") && (bottomLeft == "M" || bottomLeft == "S")
        secondDiag := tlExists && brExists && (topLeft != bottomRight) && (topLeft == "M" || topLeft == "S") && (bottomRight == "M" || bottomRight == "S")

        if firstDiag && secondDiag {
            ans += 1
        }
    }
    return ans
}

func (day4) Solve() {
    log.Println("Day 4 Solution")
    contents, err := utils.Io.ReadToString("./input/in4")
    if err != nil {
        return 
    }
    log.Printf("Part 1: %d", Day4.Part1(contents))
    log.Printf("Part 2: %d", Day4.Part2(contents))
}
