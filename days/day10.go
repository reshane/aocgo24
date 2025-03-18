package days

import (
    "log"
    "strings"

    "github.com/reshane/aocgo24/utils"
)

type day10 struct{}
var Day10 day10

func (day10) Solve() {
    log.Printf("Day 10 Solutions:\n")
    contents, err := utils.Io.ReadToString("./input/in10")
    if err != nil {
        return
    }
    log.Printf("Part 1: %d\n", Day10.Part1(contents))
    log.Printf("Part 2: %d\n", Day10.Part2(contents))
}

func (day10) parseInput(contents string) (map[Pos]int, []Pos) {
    topo := make(map[Pos]int)
    trailheads := make([]Pos, 0)
    for y, line := range strings.Split(contents, "\n") {
        if len(line) == 0 {
            continue
        }
        for x, c := range line {
            el := int(c - 48)
            if el < 0 || el > 9 {
                el = -1
            }
            if el == 0 {
                trailheads = append(trailheads, Pos{ x, y })
            }
            topo[Pos{ x, y }] = el
        }
    }
    return topo, trailheads
}

func (day10) Part1(contents string) int {
    dirs := []Pos{ { 0, -1 }, { 0, 1 }, { -1, 0 }, { 1, 0 } }
    topo, trailheads := Day10.parseInput(contents)
    // bfs from each trailhead
    // if a 9 is encountered, increment trailhead's score
    var ans int = 0
    for _, trailhead := range trailheads {
        queue := make([]Pos, 0)
        queue = append(queue, trailhead)
        peaks := utils.NewSet[Pos]()
        visited := utils.NewSet[Pos]()
        for {
            // while there is something on the queue
            // pop the first element
            curr := queue[0]
            queue = queue[1:]

            // log.Println(curr)
            // log.Println(queue)

            // append all neighbors if they are within 1 of curr
            currEl, exists := topo[curr]
            if !exists {
                log.Fatalf("%v does not exist but made it onto the queue!\n", curr)
            }

            visited.Push(curr)

            for i := 0; i < len(dirs); i++ {
                nbor := curr.plus(dirs[i])
                nborEl, exists := topo[nbor]
                if exists && nborEl == currEl + 1 {
                    if nborEl == 9 {
                        peaks.Push(nbor)
                    } else if !visited.Contains(nbor) {
                        queue = append(queue, nbor)
                    }
                }
            }

            if len(queue) == 0 {
                break
            }
        }
        ans += peaks.Size()
    }

    return ans
}

func (day10) Part2(contents string) int {
    dirs := []Pos{ { 0, -1 }, { 0, 1 }, { -1, 0 }, { 1, 0 } }
    topo, trailheads := Day10.parseInput(contents)
    // bfs from each trailhead
    // if a 9 is encountered, increment trailhead's score
    var ans int = 0
    for _, trailhead := range trailheads {
        queue := make([]Pos, 0)
        queue = append(queue, trailhead)
        visited := utils.NewSet[Pos]()
        for {
            // while there is something on the queue
            // pop the first element
            curr := queue[0]
            queue = queue[1:]

            // log.Println(curr)
            // log.Println(queue)

            // append all neighbors if they are within 1 of curr
            currEl, exists := topo[curr]
            if !exists {
                log.Fatalf("%v does not exist but made it onto the queue!\n", curr)
            }

            visited.Push(curr)

            for i := 0; i < len(dirs); i++ {
                nbor := curr.plus(dirs[i])
                nborEl, exists := topo[nbor]
                if exists && nborEl == currEl + 1 {
                    if nborEl == 9 {
                        ans += 1
                    } else if !visited.Contains(nbor) {
                        queue = append(queue, nbor)
                    }
                }
            }

            if len(queue) == 0 {
                break
            }
        }
    }

    return ans
}
