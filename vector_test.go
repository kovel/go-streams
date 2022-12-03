package main

import (
	"github.com/tj/assert"
	"testing"
)

func TestPushAndToSlice(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, []Int{1, 2, 3, 100, 1}, vec.ToSlice())
}

func TestIndexOf(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, 3, vec.IndexOf(100))
}

func TestLastIndexOf(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, 4, vec.LastIndexOf(1))
}

func TestRetainAll(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.True(t, vec.RetainAll(Vector[Int]{1, 2, 3}))
	assert.False(t, vec.RetainAll(Vector[Int]{1, 2, 3, 4}))
}

func TestDelete(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, []Int{1, 3, 100, 1}, vec.Delete(1).ToSlice())
}

func TestContains(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.True(t, vec.Contains(1))
	assert.True(t, vec.Contains(100))
	assert.False(t, vec.Contains(1100))
}

func TestGetOrDefault(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, Int(3), vec.GetOrDefault(2, 1000))
	assert.Equal(t, Int(1000), vec.GetOrDefault(20, 1000))
}

func TestLenAndIsEmpty(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, len(vec.ToSlice()), vec.Len())
	assert.False(t, vec.IsEmpty())

	vec.Clear()
	assert.True(t, vec.IsEmpty())
}

func TestAddVector(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	vec.AddVector(Vector[Int]{900, 9000})
	assert.Equal(t, []Int{1, 2, 3, 100, 1, 900, 9000}, vec.ToSlice())
}

func TestAddArray(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	vec.AddArray([]Int{9002, 90002})
	assert.Equal(t, []Int{1, 2, 3, 100, 1, 9002, 90002}, vec.ToSlice())
}

func TestAddArgs(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	vec.AddArgs(9003, 90003)
	assert.Equal(t, []Int{1, 2, 3, 100, 1, 9003, 90003}, vec.ToSlice())
}

func TestDeleteElementAndElements(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, []Int{1, 2, 100, 1}, vec.DeleteElement(3).ToSlice())
	assert.Equal(t, []Int{100}, vec.DeleteAllElements(Vector[Int]{1, 2}).ToSlice())
}

func TestAddAt(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, []Int{1, 2, 3, 200, 100, 1}, vec.AddAt(3, 200).ToSlice())
}

func TestClone(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1)
	assert.Equal(t, vec.ToSlice(), vec.Clone().ToSlice())
	assert.NotEqual(t, vec.ToSlice(), vec.Clone().Delete(0).ToSlice())
}

func TestPop(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	assert.Equal(t, Int(100), vec.Clone().Pop())
}

func TestForEach(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.ForEach(func(v Int, index int) {
		assert.Equal(t, vec.ToSlice()[index], v)
	})
}

func TestDeleteIf(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.DeleteIf(func(v Int, index int) bool {
		return v%2 == 0
	})
	assert.Equal(t, []Int{1, 3}, vec.ToSlice())
}

func TestRemoveRange(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.RemoveRange(1, 3)
	assert.Equal(t, []Int{1, 100}, vec.ToSlice())
}

func TestAddVectorAt(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.AddVectorAt(1, Vector[Int]{100, 200, 300})
	assert.Equal(t, []Int{1, 100, 200, 300, 2, 3, 100}, vec.ToSlice())
}

func TestSet(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.Set(1, 1000)
	assert.Equal(t, []Int{1, 1000, 3, 100}, vec.ToSlice())
}

func TestReplaceAll(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100)
	vec.ReplaceAll(func(v Int) Int { return v * 2 })
	assert.Equal(t, []Int{2, 4, 6, 200}, vec.ToSlice())
}

func TestSubSlice(t *testing.T) {
	vec := Vector[Int]{}
	vec.Push(1).Push(2).Push(3).Push(100).Push(1000)
	assert.Equal(t, []Int{3, 100, 1000}, vec.SubSlice(2, vec.Len()).ToSlice())
}
