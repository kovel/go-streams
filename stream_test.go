package main

import (
	"fmt"
	"github.com/tj/assert"
	"math/rand"
	"testing"
)

func TestStreamOf(t *testing.T) {
	stream := StreamOf[Int](1, 2, 2, 3, 4, 3, 100)
	assert.Equal(t, []Int{1, 2, 2, 3, 4, 3, 100}, stream.slice)
}

func TestDistinct(t *testing.T) {
	stream := StreamOf[Int](1, 2, 2, 3, 4, 3, 100)
	stream = stream.Distinct()
	stream = stream.Sort(func(i, j Int) bool {
		return i < j
	})
	assert.Equal(t, []Int{1, 2, 3, 4, 100}, stream.slice)
}

func TestFilter(t *testing.T) {
	stream := StreamOf[Int](1, 2, 2, 3, 4, 3, 100)
	stream = stream.Filter(func(v Int, i int) bool { return i > 2 })
	stream = stream.Sort(func(i, j Int) bool {
		return i < j
	})
	assert.Equal(t, []Int{3, 3, 4, 100}, stream.slice)
}

func TestReduce(t *testing.T) {
	stream := StreamOf[Int](1, 2, 3)
	assert.Equal(t, Int(6), stream.Reduce(func(a, b Int) Int { return a + b }))
}

func TestMax(t *testing.T) {
	stream := StreamOf[Int](1, 2000, 2, 3)
	assert.Equal(t, Int(2000), stream.Max(func(a, b Int) int { return int(a) - int(b) }))
}

func TestMin(t *testing.T) {
	stream := StreamOf[Int](1, 2000, 2, 3)
	assert.Equal(t, Int(1), stream.Min(func(a, b Int) int { return int(a) - int(b) }))
}

func TestNoneMatch(t *testing.T) {
	stream := StreamOf[Int](1, 2, 3, 4)
	assert.False(t, stream.NoneMatch(func(v Int) bool { return v > -1 }))

	stream2 := StreamOf[Int](1, 2, 3, 4).Map(func(v Int, i int) Int { return Int(int(v) * -1) })
	assert.True(t, stream2.NoneMatch(func(v Int) bool { return v > -1 }))
}

func TestFindAnyAndFindFirst(t *testing.T) {
	stream := StreamOf[Int](1, 2, 2, 3, 4, 3, 100)
	assert.Equal(t, stream.slice[0], stream.FindAny())
	assert.Equal(t, stream.slice[0], stream.FindFirst())
}

func TestSkipAndLimit(t *testing.T) {
	stream := StreamOf[Int](1, 2, 2, 3, 4, 3, 100)
	assert.Equal(t, stream.slice[1:3], stream.Skip(1).Limit(2).slice)
}

func TestStreamFlatMap(t *testing.T) {
	humans := Vector[Human]{}
	for i := 0; i < 10; i++ {
		pets := &Vector[Pet]{}
		for j := 0; j < rand.Intn(20); j++ {
			pets.Push(Pet{Type: fmt.Sprintf("Type %d %d\n", i, j)})
		}
		humans.Push(Human{Name: fmt.Sprintf("Name #%d", rand.Intn(100)), Pets: pets})
	}

	flatPetsStream := StreamFlatMap[Human, Pet](humans.Stream(), func(v Human, i int) *Stream[Pet] {
		return v.Pets.Stream()
	})
	assert.Equal(t, flatPetsStream.slice[0].Type, fmt.Sprintf("Type %d %d\n", 0, 0))
}
