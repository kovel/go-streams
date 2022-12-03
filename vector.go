package main

import (
	"errors"
	"sort"
	"strings"
)

type Vector[T CollectionInterface] []T

func (v *Vector[T]) ToSlice() []T {
	return *v
}

func (v *Vector[T]) Stream() *Stream[T] {
	return &Stream[T]{v.Clone().ToSlice()}
}

func (v *Vector[T]) AddVectorAt(index int, other Vector[T]) *Vector[T] {
	newV := v.Clone().ToSlice()[:index]
	newV = append(newV, other.ToSlice()...)
	newV = append(newV, v.Clone().ToSlice()[index:]...)
	*v = newV
	return v
}

func (v *Vector[T]) AddVector(other Vector[T]) *Vector[T] {
	*v = append(v.ToSlice(), other.ToSlice()...)
	return v
}

func (v *Vector[T]) AddArray(other []T) *Vector[T] {
	*v = append(v.ToSlice(), other...)
	return v
}

func (v *Vector[T]) AddArgs(other ...T) *Vector[T] {
	*v = append(v.ToSlice(), other...)
	return v
}

func (v *Vector[T]) RetainAll(other Vector[T]) bool {
	for _, value := range other.ToSlice() {
		if !v.Contains(value) {
			return false
		}
	}
	return true
}

func (v *Vector[T]) ReplaceAll(mapper func(T) T) *Vector[T] {
	var newV []T
	for _, value := range v.ToSlice() {
		newV = append(newV, mapper(value))
	}
	*v = newV
	return v
}

func (v *Vector[T]) Sort(less func(T, T) bool) *Vector[T] {
	vSlice := v.ToSlice()
	sort.Slice(vSlice, func(i, j int) bool {
		return less(vSlice[i], vSlice[j])
	})
	return v
}

func (v *Vector[T]) SubSlice(i, j int) *Vector[T] {
	newV := &Vector[T]{}
	newV.AddVector(v.ToSlice()[i:j])
	return newV
}

func (v *Vector[T]) Clear() {
	*v = make([]T, 0)
}

func (v *Vector[T]) Clone() *Vector[T] {
	newV := &Vector[T]{}
	for _, value := range v.ToSlice() {
		newV.Push(value)
	}
	return newV
}

func (v *Vector[T]) ForEach(f func(T, int)) {
	for i, t := range v.ToSlice() {
		f(t, i)
	}
}

func (v *Vector[T]) Push(value T) *Vector[T] {
	*v = append(v.ToSlice(), value)
	return v
}

func (v *Vector[T]) Pop() T {
	if len(*v) == 0 {
		var defValue T
		return defValue
	}

	last := (*v)[len(*v)-1]
	newV := (*v)[:len(*v)-1]
	*v = newV
	return last
}

func (v *Vector[T]) AddAt(index int, value T) *Vector[T] {
	var newV []T
	newV = append(newV, v.ToSlice()[:index]...)
	newV = append(newV, value)
	newV = append(newV, v.ToSlice()[index:]...)
	*v = newV
	return v
}

func (v *Vector[T]) Set(index int, value T) *Vector[T] {
	var newV []T
	newV = append(newV, v.ToSlice()[:index]...)
	newV = append(newV, value)
	newV = append(newV, v.ToSlice()[index+1:]...)
	*v = newV
	return v
}

func (v *Vector[T]) Get(index int) (T, error) {
	if index >= v.Len() {
		var defaultValue T
		return defaultValue, errors.New("index not found")
	}
	return (*v)[index], nil
}

func (v *Vector[T]) Len() int {
	return len(*v)
}

func (v *Vector[T]) IsEmpty() bool {
	return v.Len() == 0
}

func (v *Vector[T]) GetOrDefault(index int, def T) T {
	if index >= len(*v) {
		return def
	}
	return (*v)[index]
}

func (v *Vector[T]) Delete(index int) *Vector[T] {
	newV := Vector[T]{}
	for i, value := range v.ToSlice() {
		if i == index {
			continue
		}
		newV = append(newV.ToSlice(), value)
	}
	*v = newV
	return v
}

func (v *Vector[T]) DeleteElement(e T) *Vector[T] {
	for index := v.IndexOf(e); index != -1; index = v.IndexOf(e) {
		v.Delete(index)
	}
	return v
}

func (v *Vector[T]) DeleteAllElements(forDelete Vector[T]) *Vector[T] {
	for _, e := range forDelete.ToSlice() {
		v.DeleteElement(e)
	}
	return v
}

func (v *Vector[T]) DeleteAll(forDelete []int) *Vector[T] {
	forDeleteMap := make(map[int]interface{})
	for _, i := range forDelete {
		forDeleteMap[i] = nil
	}

	newV := Vector[T]{}
	for i, value := range v.ToSlice() {
		if _, ok := forDeleteMap[i]; ok {
			continue
		}
		newV = append(newV.ToSlice(), value)
	}
	*v = newV
	return v
}

func (v *Vector[T]) RemoveRange(from, to int) *Vector[T] {
	var newV []T
	newV = append(newV, (*v)[:from]...)
	newV = append(newV, (*v)[to:]...)
	*v = newV
	return v
}

func (v *Vector[T]) DeleteIf(predicate func(T, int) bool) bool {
	var forDelete []int
	for i, t := range v.ToSlice() {
		if predicate(t, i) {
			forDelete = append(forDelete, i)
		}
	}

	v.DeleteAll(forDelete)

	return len(forDelete) > 0
}

func (v *Vector[T]) IndexOf(value T) int {
	for i, item := range v.ToSlice() {
		if item == value {
			return i
		}
	}
	return -1
}

func (v *Vector[T]) LastIndexOf(value T) int {
	lastIndex := -1
	for i, item := range v.ToSlice() {
		if item == value {
			lastIndex = i
		}
	}
	return lastIndex
}

func (v *Vector[T]) Contains(value T) bool {
	return v.IndexOf(value) != -1
}

func (v *Vector[T]) String() string {
	var ret []string
	for _, v := range v.ToSlice() {
		ret = append(ret, v.String())
	}
	return strings.Join(ret, ", ")
}
