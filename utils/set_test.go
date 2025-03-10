package utils

import (
    "testing"
)

func Test_Intersection(t *testing.T) {
    a := NewSetFrom([]int{ 0, 1, 2, 3, 4 })
    b := NewSetFrom([]int{ 4, 5, 6, 7, 8 })

    c := a.Intersect(b)
    if !c.Contains(4) {
        t.Errorf("Intersection should contain only [%d]", 4)
    }
    if c.Size() != 1 {
        t.Errorf("Intersection should contain only [%d]", 4)
    }
}

func Test_Union(t *testing.T) {
    a := NewSetFrom([]int{ 0, 1, 2, 3, 4 })
    b := NewSetFrom([]int{ 4, 5, 6, 7, 8 })

    c := a.Union(b)
    for _, e := range a.Iterate() {
        if !c.Contains(e) {
            t.Errorf("Union should contain [%d]", e)
        }
    }
    for _, e := range b.Iterate() {
        if !c.Contains(e) {
            t.Errorf("Union should contain [%d]", e)
        }
    }
    if c.Size() != 9 {
        t.Errorf("Intersection should contain [%d] elements", 9)
    }
}

func Test_Difference(t *testing.T) {
    a := NewSetFrom([]int{ 0, 1, 2, 3, 4 })
    b := NewSetFrom([]int{ 4, 5, 6, 7, 8 })

    c := a.Difference(b)
    if c.Contains(4) {
        t.Errorf("Union should not contain [%d]", 4)
    }
    a.Pop(4)
    for _, e := range a.Iterate() {
        if !c.Contains(e) {
            t.Errorf("Union should contain [%d]", e)
        }
    }
    if c.Size() != 4 {
        t.Errorf("Intersection should contain [%d] elements", 4)
    }
}

func Test_Clone(t *testing.T) {
    a := NewSetFrom([]int{ 0 })
    b := a.Clone()

    b.Pop(0)
    if !a.Contains(0) {
        t.Errorf("Cloned set should contain [%d]", 0)
    }
}
