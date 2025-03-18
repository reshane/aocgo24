package days

import (
    "log"

    "github.com/reshane/aocgo24/utils"
)

type day9 struct{}
var Day9 day9

func (day9) Solve() {
    log.Printf("Day 9 Solutions:\n")

    contents, err := utils.Io.ReadToString("./input/in9")
    if err != nil {
        return
    }

    log.Printf("Part 1: %d\n", Day9.Part1(contents))
    log.Printf("Part 2: %d\n", Day9.Part2(contents))
}

func (day9) parseInput(contents string) ([]int, []Pos) {
    diskMap := make([]int, 0)
    blocks := make([]Pos, 0)
    idx := 0
    for i, c := range contents {
        size := int(c - 48)
        if size < 0 {
            continue
        }
        if i % 2 == 0 {
            blocks = append(blocks, Pos{ idx, size })
            for j := idx; j < idx + size; j++ {
                diskMap = append(diskMap, idx)
            }
            idx += 1
        } else {
            for j := idx; j < idx + size; j++ {
                diskMap = append(diskMap, -1)
            }
        }
    }
    return diskMap, blocks
}

func (day9) Part1(contents string) int {
    diskMap, _ := Day9.parseInput(contents)
    for i := len(diskMap)-1; i > -1; i-- {
        if diskMap[i] == -1 {
            continue
        }

        for j := 0; j < len(diskMap) && j < i; j++ {
            if diskMap[j] == -1 {
                diskMap[j] = diskMap[i]
                diskMap[i] = -1
            }
        }
    }
    var ans int = 0
    for i := 0; i < len(diskMap) && diskMap[i] != -1; i++ {
        ans += i * diskMap[i]
    }
    return ans
}

func insert(diskMap []int, block Pos) []int {
    // find idx that works
    inserteableIdx := -1
    for i := 0; i < len(diskMap) - block.y; i++ {
        if diskMap[i] == block.x {
            return diskMap
        }
        chunk := diskMap[i:i + block.y]
        valid := true
        for _, e := range chunk { if e != -1 { valid = false; break; } }
        if valid {
            inserteableIdx = i
            break
        }
    }
    if inserteableIdx == -1 {
        return diskMap
    }
    // if found, place the block
    for i := inserteableIdx; i < inserteableIdx + block.y; i++ {
        diskMap[i] = block.x
    }
    // remove block from old position
    for i := inserteableIdx + block.y; i < len(diskMap); i++ {
        if diskMap[i] == block.x {
            diskMap[i] = -1
        }
    }
    return diskMap
}

func (day9) Part2(contents string) int {
    diskMap, blocks := Day9.parseInput(contents)
    for i := len(blocks)-1; i > -1; i-- {
        diskMap = insert(diskMap, blocks[i])
    }
    var ans int = 0
    for i := 0; i < len(diskMap); i++ {
        if diskMap[i] != -1 {
            ans += i * diskMap[i]
        }
    }
    return ans
}
