package days

import (
    "log"
    "strings"
    "slices"

    "github.com/reshane/aocgo24/utils"
)

type day12 struct{}
var Day12 day12

func (day12) Solve() {
    log.Printf("Day 12 Solutions:\n")
    contents, err := utils.Io.ReadToString("./input/in12")
    if err != nil {
        return
    }
    log.Printf("Part 1: %d\n", Day12.Part1(contents))
    log.Printf("Part 2: %d\n", Day12.Part2(contents))
}

func (day12) parseInput(contents string) map[Pos]string {
    fences := map[Pos]string{}
    lines := strings.Split(contents, "\n")
    for y, line := range lines {
        if len(line) == 0 { continue }

        // put each pos in with its letter
        for x, _ := range line {
            fences[Pos{x, y}] = line[x:x+1]
        }
    }
    return fences
}

func (day12) Part1(contents string) int {
    dirs := []Pos{ {-1,0}, {1,0}, {0,-1}, {0,1} }
    var total int = 0
    fences := Day12.parseInput(contents)
    visitedGlobal := utils.NewSet[Pos]()
    for p, n := range fences {
        if visitedGlobal.Contains(p) { continue }
        queue := []Pos{p}
        sectionArea := 0
        sectionPerim := 0
        for {
            // log.Println(queue)
            if len(queue) == 0 { break }
            curr := queue[0]
            queue = queue[1:]

            visitedGlobal.Push(curr)

            sectionArea += 1
            nodePerim := 4
            for _, dir := range dirs {
                if nbor, exists := fences[curr.plus(dir)]; exists && nbor == n {
                    nodePerim -= 1
                    if visitedGlobal.Contains(curr.plus(dir)) { continue }
                    if slices.Contains(queue, curr.plus(dir)) { continue }
                    queue = append(queue, curr.plus(dir))
                }
            }
            sectionPerim += nodePerim
        }
        total += sectionArea * sectionPerim
    }
    return total
}

func (day12) findCorners(curr Pos, nbors []Pos, fences map[Pos]string) int {
    diag := []Pos{ {-1,-1}, {1,-1}, {-1,1}, {1,1} }
    n, exists := fences[curr]
    if !exists {
        return 0
    }
    if len(nbors) == 0 { return 4 }
    if len(nbors) == 1 { return 2 }
    if len(nbors) == 2 {
        // XXX -> 0 corners
        // 
        // X
        // X
        // X   -> 0 corners
        // 
        // X
        // XX  -> 2 corners
        // 
        // XX
        // XX  -> 1 corner

        if nbors[0].x == nbors[1].x || nbors[0].y == nbors[1].y { return 0 }

        d1 := Pos{nbors[0].x, nbors[1].y}
        d2 := Pos{nbors[1].x, nbors[0].y}
        if d1n, exists := fences[d1]; exists && d1 != curr {
            // empty diagonal
            if d1n == n { return 1 }
            return 2
        }
        if d2n, exists := fences[d2]; exists && d2 != curr {
            if d2n == n { return 1 }
            return 2
        }
    }
    if len(nbors) == 3 {
        // T configuration
        // 
        // XXX
        //  X
        // 
        // X
        // XX
        // X

        sectionCorners := 0
        if nbors[0].x == nbors[1].x {
            // 0 and 1 colinear in y direction
            // take the y of 0 and the x of 2
            // take the y of 1 and the x of 2
            if nbor, exists := fences[Pos{nbors[2].x, nbors[0].y}]; exists && nbor != n && (Pos{nbors[2].x, nbors[0].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[2].x, nbors[1].y}]; exists && nbor != n && (Pos{nbors[2].x, nbors[1].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
        if nbors[0].y == nbors[1].y {
            // 0 and 1 colinear in the x direction
            // take the x of 0 and the y of 2
            // take the x of 1 and the y of 2
            if nbor, exists := fences[Pos{nbors[0].x, nbors[2].y}]; exists && nbor != n && (Pos{nbors[0].x, nbors[2].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[1].x, nbors[2].y}]; exists && nbor != n && (Pos{nbors[1].x, nbors[2].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
        if nbors[2].x == nbors[1].x {
            if nbor, exists := fences[Pos{nbors[0].x, nbors[2].y}]; exists && nbor != n && (Pos{nbors[0].x, nbors[2].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[0].x, nbors[1].y}]; exists && nbor != n && (Pos{nbors[0].x, nbors[1].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
        if nbors[2].y == nbors[1].y {
            if nbor, exists := fences[Pos{nbors[2].x, nbors[0].y}]; exists && nbor != n && (Pos{nbors[2].x, nbors[0].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[1].x, nbors[0].y}]; exists && nbor != n && (Pos{nbors[1].x, nbors[0].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
        if nbors[2].x == nbors[0].x {
            if nbor, exists := fences[Pos{nbors[1].x, nbors[2].y}]; exists && nbor != n && (Pos{nbors[1].x, nbors[2].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[1].x, nbors[0].y}]; exists && nbor != n && (Pos{nbors[1].x, nbors[0].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
        if nbors[2].y == nbors[0].y {
            if nbor, exists := fences[Pos{nbors[2].x, nbors[1].y}]; exists && nbor != n && (Pos{nbors[2].x, nbors[1].y}) != curr { sectionCorners++ }
            if nbor, exists := fences[Pos{nbors[0].x, nbors[1].y}]; exists && nbor != n && (Pos{nbors[0].x, nbors[1].y}) != curr { sectionCorners++ }
            return sectionCorners
        }
    }
    // check corners
    // number of missing corners is numer of corners
    dNbors := 0
    for _, dir := range diag {
        if d, exists := fences[curr.plus(dir)]; exists && d == n {
            dNbors++
        }
    }
    return 4 - dNbors
}

func (day12) Part2(contents string) int {
    dirs := []Pos{ {-1,0}, {1,0}, {0,-1}, {0,1} }
    // find corners for each mass
    var total int = 0
    fences := Day12.parseInput(contents)
    visitedGlobal := utils.NewSet[Pos]()
    for p, n := range fences {
        if visitedGlobal.Contains(p) { continue }
        queue := []Pos{p}
        sectionCorners := 0
        sectionArea := 0
        for {
            if len(queue) == 0 { break }
            curr := queue[0]
            queue = queue[1:]

            visitedGlobal.Push(curr)
            sectionArea += 1

            nbors := []Pos{}
            for _, dir := range dirs {
                if nbor, exists := fences[curr.plus(dir)]; exists && nbor == n {
                    nbors = append(nbors, curr.plus(dir))
                    if visitedGlobal.Contains(curr.plus(dir)) { continue }
                    if slices.Contains(queue, curr.plus(dir)) { continue }
                    queue = append(queue, curr.plus(dir))
                }
            }

            corners := Day12.findCorners(curr, nbors, fences)
            sectionCorners += corners
        }
        total += sectionArea * sectionCorners
    }
    return total
}
