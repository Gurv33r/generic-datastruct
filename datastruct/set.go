package datastruct

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	elements []T
}

//constructors
func NewSet[T comparable]() (set *Set[T]) {
	set = new(Set[T])
	return
}

func NewSetFrom[T comparable](elements ...T) (set *Set[T]) {
	set = new(Set[T])
	for _, v := range elements {
		set.Add(v)
	}
	return
}

//accessors and mutators
func (set *Set[T]) Length() int {
	return len(set.elements)
}

func (set *Set[T]) Add(elements ...T) bool {
	for _, v := range elements {
		if set.Contains(v) {
			return false
		}
		set.elements = append(set.elements, v)
	}
	return true
}

func (set *Set[T]) Contains(element T) bool {
	for _, v := range set.elements {
		if v == element {
			return true
		}
	}
	return false
}

func (set *Set[T]) Elements() []T {
	return set.elements
}

func (set *Set[T]) Remove(element T) bool {
	if !set.Contains(element) {
		return false
	}
	if set.elements[0] == element {
		set.elements = set.elements[1:]
	} else if set.elements[set.Length()-1] == element {
		set.elements = set.elements[:set.Length()-1]
	} else {
		for i, v := range set.elements {
			if v == element {
				set.elements = append(set.elements[:i], set.elements[i+1:]...)
			}
		}
	}
	return true
}

//helper functions
func (set *Set[T]) String() string {
	var bldr strings.Builder
	bldr.WriteRune('{')
	for i, elem := range set.elements {
		bldr.WriteString(fmt.Sprintf("%v", elem))
		if i != len(set.elements)-1 {
			bldr.WriteString(", ")
		}
	}
	bldr.WriteRune('}')
	return bldr.String()
}
