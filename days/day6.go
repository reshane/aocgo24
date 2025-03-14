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
    guard, walls := Day6.parseInput(contents)
    visited := utils.NewSet[Vel]()
    path := make([]Vel, 0)
    dirs := []Pos{ {0, -1}, {1, 0}, {0, 1}, {-1, 0} }
    dirIdx := 0
    curr := guard
    for {
        path = append(path, Vel{curr, dirIdx})
        visited.Push(Vel{curr, dirIdx})

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
    // for every position in path
    // place obstacle
    // step until either pos and dir is in visited
    // or pos does not exists
    checked := utils.NewSet[Pos]()
    cycles := utils.NewSet[Pos]()
    checked.Push(guard)
    for i := 0; i < len(path)-1; i++ {
        if checked.Contains(path[i+1].pos) {
            // if the obstacle position is in the path before the curr pos,
            // then the curr pos might be unreachable from here
            // and we have also already checked it
            continue
        }
        curr = guard
        dirIdx = 0
        walls[path[i+1].pos] = 1
        cycleVisited := utils.NewSet[Vel]()
        // check if this causes a cycle
        for {
            cycleVisited.Push(Vel{ curr, dirIdx })
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

            if cycleVisited.Contains(Vel{ curr, dirIdx }) {
                // cycle!!
                cycles.Push(path[i+1].pos)
                break
            }
        }
        walls[path[i+1].pos] = 0
    }
    return cycles.Size()
}
