package days

import (
    "log"
    "strings"
    "strconv"
    "sort"
    "slices"

    "github.com/reshane/aocgo24/utils"
)

type day5 struct{}
var Day5 day5

func parseInput(contents string) (map[int][]int, [][]int) {
    parts := strings.Split(contents, "\n\n")
    mappings := strings.Split(parts[0], "\n")
    pages := strings.Split(parts[1], "\n")

    topo := make(map[int][]int)
    for i := 0; i < len(mappings); i++ {
        split := strings.Split(mappings[i], "|")
        first, err := strconv.Atoi(split[0])
        if err != nil {
            log.Fatalf("first Could not convert %s to int", split[0])
        }
        second, err := strconv.Atoi(split[1])
        if err != nil {
            log.Fatalf("second Could not convert %s to int", split[1])
        }

        val, exists := topo[first]
        if exists {
            topo[first] = append(val, second)
        } else {
            topo[first] = append(make([]int, 0), second)
        }
    }

    orderings := make([][]int, 0)
    for _, ordering := range pages {
        if len(strings.TrimSpace(ordering)) < 1 {
            continue
        }
        order := strings.Split(ordering, ",")
        pgs := make([]int, 0, len(order))
        for _, page := range order {
            if len(page) < 1 {
                continue
            }
            pageInt, err := strconv.Atoi(page)
            if err != nil {
                log.Fatalf("ordering Could not convert %s to int", page)
            }
            pgs = append(pgs, pageInt)
        }
        orderings = append(orderings, pgs)
    }
    return topo, orderings
}

func checkTopo(topo map[int][]int, order []int) bool {
    set := make(map[int]struct{})
    orderingSet := make(map[int]struct{})
    for _, o := range order {
        orderingSet[o] = struct{}{}
    }
    for i := 0; i < len(order); i++ {
        curr := order[i]
        deps := topo[curr]
        for _, dep := range deps {
            if _, e := orderingSet[dep]; e {
                set[dep] = struct{}{}
            }
        }
        delete(set, curr)
    }
    return len(set) == 0
}

func (day5) Part1(contents string) int {
    topo, orderings := parseInput(contents)
    var ans int = 0
    for _, ordering := range orderings {
        if checkTopo(topo, ordering) {
            ans += ordering[len(ordering) / 2]
        }
    }
    return ans
}

func (day5) Part2(contents string) int {
    topo, orderings := parseInput(contents)
    var ans int = 0
    for _, ordering := range orderings {
        if !checkTopo(topo, ordering) {
            sort.Slice(ordering, func(i, j int) bool { 
                deps := topo[ordering[i]]
                return slices.Contains(deps, ordering[j])
            })
            ans += ordering[len(ordering) / 2]
        }
    }
    return ans
}

func (day5) Solve() {
    log.Println("Day 5 Solution")
    contents, err := utils.Io.ReadToString("./input/in5")
    if err != nil {
        return
    }

    log.Printf("Part 1: %d", Day5.Part1(contents))
    log.Printf("Part 2: %d", Day5.Part2(contents))
}
