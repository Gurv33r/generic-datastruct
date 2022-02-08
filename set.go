package datastruct

type Set[T comparable] struct{
	elements []T
}

func NewSet[T comparable]() (set *Set[T]) {
	set = new(Set[T])
	return
}
func (set *Set[T]) Length() int{
	return len(set.elements)
}

func (set *Set[T]) Add(elements ...T) bool{
	for _, v := range elements{
		if set.Contains(element){
			return false
		}
		set.elements = append(set.elements, element)
	}
	return true
}

func (set *Set[T]) Contains(element T) bool{
	for _, v := range set.elements{
		if v == element{
			return true 
		}
	}
	return false
}

func (set *Set[T]) Elements() []T {
	return set.elements
}
func NewSetFrom[T comparable] (elements ...T) (set *Set[T]) {
	set = new(Set[T])
	for _, v := range elements{
		set.Add(v)
	}
	return
}
func (set *Set[T]) Remove(element T) bool {
	if !set.Contains(element){
		return false
	}
	if set.elements[0] == element{
		set.elements = set.elements[1:]
	} else if set.elements[set.Length()-1] == element{
		set.elements = set.elements[:set.Length()-1]
	} else {
		for i, v := range set.elements{
			if v == element {
				set.elements = append(set.elements[:i], set.elements[i+1:]...)
			}
		}
	}
	return true 
}