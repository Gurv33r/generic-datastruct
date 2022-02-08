package datastruct_test

import "github.com/Gurv33r/generic-datastruct"

import "testing"

func TestNewSet(t *testing.T) {
	set := datastruct.NewSet[int]()
	elems := set.Elements()
	switch v := elems.(type) {
	case int:
		t.Skip()
	default:
		t.Errorf("Set not of right type:\n\tExpected = %T\n\tactual = %T\n", []int{}, elems)
	}
}

func TestNewSetFrom(t *testing.T) {
	set := datastruct.NewSetFrom[int](1, 2, 3, 4, 1)
	actual := []int{1, 2, 3, 4}
	for i, v := range set.Elements {
		if v != actual[i] {
			t.Errorf("Inaccurate data value stored:\n\tExpected = %d\n\tActual = %d\n", actual[i], v)
		}
	}
}

func TestLength(t *testing.T) {
	set1 := datastruct.NewSet[any]()
	if set1.Length() != 0 {
		t.Errorf("Empty set declaration isn't empty!\n\tExpected set length = %d\n\tActual set length = %d\n", 0, set1.Length())
	}
	set1.Add('A')
	if set1.Length() != 1 {
		t.Errorf("Adding to set doesn't update the length!\n\tExpected set length = %d\n\tActual set length = %d\n", 1, set1.Length())
	}
	set1.Remove('A')
	if set1.Length() != 0 {
		t.Errorf("Removing from set doesn't update the length!\n\tExpected set length = %d\n\tActual set length = %d\n", 0, set1.Length())
	}
}

func TestAdd(t *testing.T) {
	set := datastruct.NewSet[int]()
	set.Add(1)
	if set.Elements()[0] != 1 {
		t.Errorf("Inaccurate element addition!\n\tExpected = %d\n\tActual = %d\n", 1, set.Elements())
	}
	set.Add(1)
	if set.Length() != 1 {
		t.Errorf("Set not filtering out duplicates!\n\tExpected = %v\n\tActual=%v\n", []int{1}, set.Elements())
	}
}

func TestRemove() (t *testing.T) {
	set := datastruct.NewSetFrom[int](1, 2, 3)
	if !set.Remove(1) {
		t.Errorf("Didn't delete element found in set!\n\tExpected = %v\n\tActual = %v\n", []int{2, 3}, set.Elements())
	}
	if !set.Remove(0) {
		t.Errorf("Deleted an element not found in set!\n\tExpected = %v\n\tActual = %v\n", []int{2, 3}, set.Elements())
	}
}

func TestContains() (t *testing.T) {
	set := datastruct.NewSetFrom[int](1, 2, 3)
	if set.Contains("h") {
		t.Errorf("Found \"h\" in %v\n", set.Elements())
	}
	if !set.Contains(1) {
		t.Errorf("Didn't find 1 in %v\n", set.Elements())
	}
}
