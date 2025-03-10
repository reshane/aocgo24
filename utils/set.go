package utils

type Set[T comparable] struct {
    data map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
    return Set[T]{ make(map[T]struct{}) }
}

func NewSetFrom[T comparable](collection []T) Set[T] {
    set := NewSet[T]()
    for _, e := range collection {
        set.Push(e)
    }
    return set
}

func (s *Set[T]) Clone() Set[T] {
    return Set[T]{ s.data }
}

func (s *Set[T]) Push(e T) {
    s.data[e] = struct{}{}
}

func (s *Set[T]) Pop(e T) {
    delete(s.data, e)
}

func (s *Set[T]) Contains(e T) bool {
    _, c := s.data[e]
    return c
}

func (s *Set[T]) Size() int {
    return len(s.data)
}

func (s *Set[T]) Intersect(other Set[T]) Set[T] {
    res := NewSet[T]()
    for e, _ := range s.data {
        if other.Contains(e) {
            res.Push(e)
        }
    }
    return res
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
    res := other.Clone()
    for e, _ := range s.data {
        res.Push(e)
    }
    return res
}

func (s *Set[T]) Difference(other Set[T]) Set[T] {
    res := s.Clone()
    for e, _ := range s.data {
        if other.Contains(e) {
            res.Pop(e)
        }
    }
    return res
}

func (s *Set[T]) Iterate() []T {
    iter := make([]T, 0)
    for e, _ := range s.data {
        iter = append(iter, e)
    }
    return iter
}
