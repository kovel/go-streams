package main

import (
	"log"
	"sort"
)

type Stream[T CollectionInterface] struct {
	slice []T
}

func StreamOf[T CollectionInterface](args ...T) *Stream[T] {
	return &Stream[T]{args}
}

func (s *Stream[T]) ForEach(f func(T, int)) {
	for i, t := range s.slice {
		f(t, i)
	}
}

func (s *Stream[T]) Map(mapper func(T, int) T) *Stream[T] {
	for i, t := range s.slice {
		s.slice[i] = mapper(t, i)
	}
	return s
}

func (s *Stream[T]) Reduce(reducer func(T, T) T) T {
	arg1 := s.slice[0]

	if len(s.slice) == 1 {
		return arg1
	}

	for _, t := range s.slice[1:] {
		arg1 = reducer(arg1, t)
	}
	return arg1
}

func (s *Stream[T]) Max(comparator func(a, b T) int) T {
	ret := s.slice[0]
	if len(s.slice) == 1 {
		return ret
	}

	for _, t := range s.slice[1:] {
		if comparator(ret, t) < 0 {
			ret = t
		}
	}
	return ret
}

func (s *Stream[T]) Min(comparator func(a, b T) int) T {
	ret := s.slice[0]
	if len(s.slice) == 1 {
		return ret
	}

	for _, t := range s.slice[1:] {
		if comparator(ret, t) > 0 {
			ret = t
		}
	}
	return ret
}

func (s *Stream[T]) NoneMatch(predicate func(T) bool) bool {
	for _, t := range s.slice {
		if predicate(t) {
			return false
		}
	}
	return true
}

func StreamFlatMap[T CollectionInterface, R CollectionInterface](s *Stream[T], mapper func(T, int) *Stream[R]) *Stream[R] {
	newSlice := make([]R, 0)
	for i, t := range s.slice {
		newSlice = append(newSlice, mapper(t, i).slice...)
	}
	return StreamOf[R](newSlice...)
}

func StreamMap[T CollectionInterface, R CollectionInterface](s *Stream[T], mapper func(T, int) R) *Stream[R] {
	newSlice := make([]R, 0)
	for i, t := range s.slice {
		newSlice = append(newSlice, mapper(t, i))
	}
	return StreamOf[R](newSlice...)
}

func (s *Stream[T]) Sort(less func(T, T) bool) *Stream[T] {
	sort.Slice(s.slice[:], func(i, j int) bool {
		return less(s.slice[i], s.slice[j])
	})
	return s
}

func (s *Stream[T]) Skip(i int) *Stream[T] {
	return StreamOf[T](s.slice[i:]...)
}

func (s *Stream[T]) Limit(i int) *Stream[T] {
	return StreamOf[T](s.slice[:i]...)
}

func (s *Stream[T]) Filter(predicate func(T, int) bool) *Stream[T] {
	newSlice := []T{}
	log.Println(s.slice)
	for i, t := range s.slice {
		if predicate(t, i) {
			newSlice = append(newSlice, t)
		}
	}
	return StreamOf[T](newSlice...)
}

func (s *Stream[T]) FindAny() T {
	if len(s.slice) == 0 {
		var defaultValue T
		return defaultValue
	}
	return s.slice[0]
}

func (s *Stream[T]) FindFirst() T {
	if len(s.slice) == 0 {
		var defaultValue T
		return defaultValue
	}
	return s.slice[0]
}

func (s *Stream[T]) AllMatch(predicate func(T, int) bool) bool {
	for i, t := range s.slice {
		if !predicate(t, i) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) AnyMatch(predicate func(T, int) bool) bool {
	for i, t := range s.slice {
		if predicate(t, i) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) Len() int {
	return len(s.slice)
}

func (s *Stream[T]) Distinct() *Stream[T] {
	m := make(map[T]interface{})
	s.ForEach(func(t T, i int) {
		m[t] = nil
	})

	newSlice := []T{}
	for v, _ := range m {
		newSlice = append(newSlice, v)
	}
	return &Stream[T]{newSlice}
}
