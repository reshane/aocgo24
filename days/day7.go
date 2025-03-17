package days

import (
    "log"
    "fmt"
    "strconv"
    "strings"

    "github.com/reshane/aocgo24/utils"
)

type day7 struct{}
var Day7 day7

func (day7) Solve() {
    log.Printf("Day 7 Solutions:\n")
    contents, err := utils.Io.ReadToString("./input/in7")
    if err != nil {
        return
    }
    log.Printf("Part 1: %d\n", Day7.Part1(contents))
    log.Printf("Part 2: %d\n", Day7.Part2(contents))
}

type eq struct {
    operands []int
    ans int
}

func (day7) parseInput(contents string) []eq {
    eqs := make([]eq, 0)
    lines := strings.Split(contents, "\n")
    for _, line := range lines {
        if len(line) < 1 {
            continue
        }
        parts := strings.Split(line, ": ")
        ans, err := strconv.Atoi(parts[0])
        if err != nil {
            log.Fatalf("Could not parse ans from string %s\n", parts[0])
        }
        operands := make([]int, 0)
        operandStrings := strings.Split(parts[1], " ")
        for i := 0; i < len(operandStrings); i++ {
            operand, err := strconv.Atoi(operandStrings[i])
            if err != nil {
                log.Fatalf("Could not parse operand from string %s\n", operandStrings[i])
            }
            operands = append(operands, operand)
        }
        eqs = append(eqs, eq{operands, ans})
    }
    return eqs
}

func mul(op1 int, op2 int) int {
    return op1 * op2
}

func add(op1 int, op2 int) int {
    return op1 + op2
}

func cat(op1 int, op2 int) int {
    cat, err := strconv.Atoi(fmt.Sprintf("%d%d", op1, op2))
    if err != nil {
        log.Fatalf("Could not convert %d and %d into catted int %v\n", op1, op2, err)
    }
    return cat
}

var opFuncs map[rune]func(int, int) int = map[rune]func(int, int) int{ '+': add, '*': mul, '|': cat }

func calcResult(op rune, op1 int, op2 int) int {
    return opFuncs[op](op1, op2)
}

func (day7) solvePart(eqs []eq, ops []rune) int {
    // for each eq
    // put first operand in a hashset
    // apply every possible operation to each element in previous hashset
    // at the end, compare every entry in the hashset to the result
    var total int = 0

    for _, eq := range eqs {
        current := utils.NewSetFrom([]int{ eq.operands[0] })
        for i := 1; i < len(eq.operands); i++ {
            next := utils.NewSet[int]()
            for _, prev := range current.Iterate() {
                // apply each operand to the element and put it in next
                for _, op := range ops {
                    next.Push(calcResult(op, prev, eq.operands[i]))
                }
            }
            current = next
        }
        for _, a := range current.Iterate() {
            if a == eq.ans {
                total += eq.ans
                continue
            }
        }
    }

    return total
}

func (day7) Part1(contents string) int {
    ops := []rune{ '+', '*' }
    eqs := Day7.parseInput(contents)
    return Day7.solvePart(eqs, ops)
}

func (day7) Part2(contents string) int {
    ops := []rune{ '+', '*', '|' }
    eqs := Day7.parseInput(contents)
    return Day7.solvePart(eqs, ops)
}
