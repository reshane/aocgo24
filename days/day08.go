package days

import (
    "log"
    "strings"

    "github.com/reshane/aocgo24/utils"
)

type day8 struct{}
var Day8 day8

func (day8) Solve() {
    log.Printf("Day 8 Solutions:\n")
    contents, err := utils.Io.ReadToString("./input/in8")
    if err != nil {
        return
    }

    log.Printf("Part 1: %d\n", Day8.Part1(contents))
    log.Printf("Part 2: %d\n", Day8.Part2(contents))
}

func (day8) parseInput(contents string) (utils.Set[Pos], map[rune][]Pos) {
    antennas := make(map[rune][]Pos)
    positions := utils.NewSet[Pos]()
    for y, line := range strings.Split(contents, "\n") {
        if len(line) == 0 {
            continue
        }
        for x, p := range line {
            positions.Push(Pos{ x, y })
            if p != '.' {
                val, exists := antennas[p]
                if exists {
                    antennas[p] = append(val, Pos{ x, y })
                } else {
                    antennas[p] = []Pos{ { x, y } }
                }
            }
        }
    }
    return positions, antennas
}

func (day8) Part1(contents string) int {
    positions, antennas := Day8.parseInput(contents)
    antinodes := utils.NewSet[Pos]()
    for _, v := range antennas {
        for i := 0; i < len(v)-1; i++ {
            for j := i+1; j < len(v); j++ {
                a := v[i];
                b := v[j];
                diff := Pos{ a.x-b.x, a.y-b.y }
                ad := Pos{ a.x+diff.x, a.y+diff.y }
                bd := Pos{ b.x-diff.x, b.y-diff.y }

                if positions.Contains(ad) {
                    antinodes.Push(ad)
                }
                if positions.Contains(bd) {
                    antinodes.Push(bd)
                }
            }
        }
    }
    return antinodes.Size()
}

func (day8) Part2(contents string) int {
    positions, antennas := Day8.parseInput(contents)
    antinodes := utils.NewSet[Pos]()
    for _, v := range antennas {
        for i := 0; i < len(v)-1; i++ {
            for j := i+1; j < len(v); j++ {
                a := v[i];
                b := v[j];
                diff := Pos{ a.x-b.x, a.y-b.y }
                k := 0
                for {
                    ad := Pos{ a.x+(diff.x * k), a.y+ (diff.y * k) }
                    bd := Pos{ b.x-(diff.x * k), b.y-(diff.y * k) }
                    adi := positions.Contains(ad)
                    bdi := positions.Contains(bd)

                    if adi {
                        antinodes.Push(ad)
                    }
                    if bdi {
                        antinodes.Push(bd)
                    }

                    if !adi && !bdi {
                        break
                    }
                    k += 1
                }
            }
        }
    }
    return antinodes.Size()
}
