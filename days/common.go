package days

type Solver interface {
    Solve()
}

var Solutions = []Solver{ 
    Day1,
    Day2,
    Day3,
    Day4,
    Day5,
    Day6,
    Day7,
    Day8,
    Day9,
    Day10,
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
