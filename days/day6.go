package days

import (
    "log"
    "strings"

    "github.com/reshane/aocgo24/utils"
)

type day6 struct{}
var Day6 day6

func (day6) Solve() {
    contents, err := utils.Io.ReadToString("./input/in6")
    if err != nil {
        return
    }

    log.Printf("Day 6 Solutions\n")
    log.Printf("Part 1: %d\n", Day6.Part1(contents))
    log.Printf("Part 2: %d\n", Day6.Part2(contents))
}

type Pos struct {
    x int
    y int
}

func (a *Pos) add(other Pos) {
    a.x += other.x
    a.y += other.y
}

func (a *Pos) plus(other Pos) Pos {
    return Pos { a.x + other.x, a.y + other.y }
}

func (day6) parseInput(contents string) (Pos, map[Pos]int) {
    guard := Pos{ -1, -1 }
    walls := make(map[Pos]int)
    lines := strings.Split(contents, "\n")
    for y, line := range lines {
        if len(line) <= 0 {
            continue
        }
        for x, c := range line {
            if c == '#' {
                walls[Pos{ x, y }] = 1
            } else {
                walls[Pos{ x, y }] = 0
            }
            if c == '^' {
                guard.x = x
                guard.y = y
            }
        }
    }

    return guard, walls
}

func (day6) Part1(contents string) int {
    guard, walls := Day6.parseInput(contents)
    visited := utils.NewSet[Pos]()
    dirs := []Pos{ {0, -1}, {1, 0}, {0, 1}, {-1, 0} }
    dirIdx := 0
    curr := guard
    for {
        visited.Push(curr)
        wall, exists := walls[curr.plus(dirs[dirIdx])]
        if !exists {
            break
        }
        if wall == 1 {
            dirIdx += 1
            dirIdx %= 4
        } else {
            curr.add(dirs[dirIdx])
        }
    }
    return visited.Size()
}

type Vel struct {
    pos Pos
    dir int
}

func (day6) Part2(contents string) int {
    dirs := []Pos{ {0, -1}, {1, 0}, {0, 1}, {-1, 0} }
    guard, walls := Day6.parseInput(contents)
    curr := Vel{ guard, 0 }
    visited := utils.NewSetFrom([]Vel{ curr })
    cycleObs := utils.NewSet[Pos]()
    for {
        // add an obstacle at the next position
        // run until exit or cycle
        // if cycle, add the next position to the set
        visited.Push(curr)
        
        // place a wall at the next position
        wall, exists := walls[curr.pos.plus(dirs[curr.dir])]
        if !exists {
            break
        }
        if wall == 1 {
            // turn and continue
            curr.dir += 1
            curr.dir %= 4
        } else {
            pn := curr.pos.plus(dirs[curr.dir])
            if visited.Contains(Vel{pn,0}) || visited.Contains(Vel{pn,1}) || visited.Contains(Vel{pn,2}) || visited.Contains(Vel{pn,3}) {
                curr.pos.add(dirs[curr.dir])
                continue
            }
            // place obstacle at curr.pos.plus(dirs[curr.dir])]
            walls[curr.pos.plus(dirs[curr.dir])] = 1
            // step until something either isn't in walls or is in cycleVisited
            cycleVisited := visited.Clone()
            cycleCurr := curr
            for {
                cycleVisited.Push(cycleCurr)
                next := Vel{cycleCurr.pos.plus(dirs[cycleCurr.dir]), cycleCurr.dir}

                wall, exists := walls[next.pos]
                if !exists {
                    // no cycle here
                    break
                }
                if wall == 1 {
                    // turn 
                    cycleCurr.dir += 1
                    cycleCurr.dir %= 4
                } else {
                    // if next is in visited, cycle
                    if cycleVisited.Contains(next) {
                        cycleObs.Push(curr.pos.plus(dirs[curr.dir]))
                        break
                    }
                    // else, curr = next
                    cycleCurr.pos.add(dirs[cycleCurr.dir])
                }
            }

            // unset the wall for other iterations
            walls[curr.pos.plus(dirs[curr.dir])] = 0
            // update current position
            curr.pos.add(dirs[curr.dir])
        }
    }
    return cycleObs.Size()
}

