package readysetgo

type Set[T comparable] map[T]struct{}

// Functions implemented by Python set:
// pop()	 	Removes an element from the set
// symmetric_difference()	^	Returns a set with the symmetric differences of two sets
// symmetric_difference_update()	^=	Inserts the symmetric differences from this set and another

func New[T comparable](i ...T) Set[T] {
	s := Set[T]{}
	s.Add(i...)
	return s
}

// Add adds one or more items to the set
func (s Set[T]) Add(i ...T) {
	for _, v := range i {
		s[v] = struct{}{}
	}
}

// Clear removes all the items from the set
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Copy returns a copy of the set
func (s Set[T]) Copy() Set[T] {
	c := New[T]()
	for k := range s {
		c.Add(k)
	}
	return c
}

// Has tells if the set contains the specified item
func (s Set[T]) Has(i T) bool {
	_, ok := s[i]
	return ok
}

// Items returns all the items in the set
func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for k := range s {
		items = append(items, k)
	}
	return items
}

// Update updates the set with the union of this set and others
func (s Set[T]) Update(others ...Set[T]) {
	for _, o := range others {
		for k := range o {
			s.Add(k)
		}
	}
}

// IsSubset returns whether another set contains this set or not
func (s Set[T]) IsSubset(o Set[T]) bool {
	for k := range s {
		if !o.Has(k) {
			return false
		}
	}
	return true
}

// IsDisjoint returns whether two sets have a intersection or not
func (s Set[T]) IsDisjoint(o Set[T]) bool {
	i := Intersection(s, o)
	return len(i) == 0
}

// IsSuperset returns whether this set contains another set or not
func (s Set[T]) IsSuperset(o Set[T]) bool {
	return o.IsSubset(s)
}

// Remove removes the specified item from the set and returns true if it was found or false otherwise
func (s Set[T]) Remove(i T) bool {
	if _, ok := s[i]; ok {
		delete(s, i)
		return true
	}
	return false
}

// DifferenceUpdate removes the items in this set that are also included in other, specified set(s)
func (s Set[T]) DifferenceUpdate(o ...Set[T]) {
	for i := range Union(o...) {
		s.Remove(i)
	}
}

// IntersectionUpdate removes the items in this set that are not present in other, specified set(s)
func (s Set[T]) IntersectionUpdate(o ...Set[T]) {
	u := Union(o...)
	for i := range s {
		if !u.Has(i) {
			s.Remove(i)
		}
	}
}

// Union returns a set containing the union of sets
func Union[T comparable](s ...Set[T]) Set[T] {
	o := New[T]()
	for _, set := range s {
		for k := range set {
			o.Add(k)
		}
	}
	return o
}

// Difference returns a set containing the difference between a set and one or more other sets
func Difference[T comparable](s ...Set[T]) Set[T] {
	d := New[T]()
	for i := range s[0] {
		found := false
		for _, o := range s[1:] {
			if o.Has(i) {
				found = true
				break
			}
		}
		if !found {
			d.Add(i)
		}
	}
	return d
}

// Intersection returns a set, that is the intersection of two or more other sets
func Intersection[T comparable](s ...Set[T]) Set[T] {
	o := New[T]()
	for i := range s[0] {
		found := true
		for _, set := range s[1:] {
			if !set.Has(i) {
				found = false
				break
			}
		}
		if found {
			o.Add(i)
		}
	}
	return o
}
