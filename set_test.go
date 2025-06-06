package readysetgo_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ilmaruk/readysetgo"
)

func TestSet_Add(t *testing.T) {
	s := readysetgo.New[int]()

	// Single item
	s.Add(1)
	assertSet(t, s, []int{1})

	// Multiple items
	s.Add(1, 2, 3)
	assertSet(t, s, []int{1, 2, 3})
}

func TestSet_Clear(t *testing.T) {
	s := readysetgo.New(1, 2, 3)
	s.Clear()
	assertSet(t, s, []int{})
}

func TestSet_Copy(t *testing.T) {
	l := readysetgo.New(1, 2, 3)
	r := l.Copy()
	require.NotEqual(t, reflect.ValueOf(l).Pointer() == reflect.ValueOf(r).Pointer(), "expected different sets")
	assertSet(t, r, []int{1, 2, 3})
}

func TestSet_Has(t *testing.T) {
	s := readysetgo.New(1)
	require.True(t, s.Has(1), "expected to find 1 in set")
	require.False(t, s.Has(2), "expected not to find 2 in set")
}

func TestSer_Items(t *testing.T) {
	s := readysetgo.New(1, 2, 3)
	require.ElementsMatch(t, []int{1, 2, 3}, s.Items(), "expected items to match")
}

func TestSet_Update(t *testing.T) {
	s := readysetgo.New(1, 2, 3)
	s.Update(readysetgo.New(3, 4, 5), readysetgo.New(5, 6, 7))
	assertSet(t, s, []int{1, 2, 3, 4, 5, 6, 7})
}

func TestDifference(t *testing.T) {
	s := readysetgo.Difference(readysetgo.New(1, 2, 3, 4), readysetgo.New(3, 4, 5),
		readysetgo.New(1, 3, 6))
	assertSet(t, s, []int{2})
}

func TestSet_IsSubset(t *testing.T) {
	s := readysetgo.New(1, 2, 3)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is subset",
			other:    readysetgo.New(0, 1, 2, 3, 4),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.New(1, 2, 3),
			expected: true,
		},
		{
			name:     "is not subset",
			other:    readysetgo.New(1, 2, 4),
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, s.IsSubset(tc.other), "expected %v to be subset of %v", s, tc.other)
		})
	}
}

func TestSet_IsDisjoint(t *testing.T) {
	s := readysetgo.New(1, 2, 3)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is disjoint",
			other:    readysetgo.New(0, 4, 5),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.New(1, 2, 3),
			expected: false,
		},
		{
			name:     "is not disjoint",
			other:    readysetgo.New(3, 4, 5),
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, s.IsDisjoint(tc.other), "expected %v to be disjoint from %v", s, tc.other)
		})
	}
}

func TestSet_IsSuperset(t *testing.T) {
	s := readysetgo.New(0, 1, 2, 3, 4)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is superset",
			other:    readysetgo.New(1, 2, 3),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.New(0, 1, 2, 3, 4),
			expected: true,
		},
		{
			name:     "is not superset",
			other:    readysetgo.New(1, 2, 5),
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, s.IsSuperset(tc.other), "expected %v to be superset of %v", s, tc.other)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	s := readysetgo.New(1, 2, 3)

	// Remove existing item
	require.True(t, s.Remove(2), "expected to remove 2 from set")
	assertSet(t, s, []int{1, 3})

	// Remove non-existing item
	require.False(t, s.Remove(4), "expected not to remove 4 from set")
	assertSet(t, s, []int{1, 3})
}

func TestSet_DifferenceUpdate(t *testing.T) {
	a := readysetgo.New("apple", "banana", "cherry")
	b := readysetgo.New("google", "microsoft", "apple")
	c := readysetgo.New("cherry", "micra", "bluebird")
	a.DifferenceUpdate(b, c)
	assertSet(t, a, []string{"banana"})
}

func TestSet_IntersectionUpdate(t *testing.T) {
	x := readysetgo.New("a", "b", "c")
	y := readysetgo.New("c", "d", "e")
	z := readysetgo.New("f", "g", "c")
	x.IntersectionUpdate(y, z)
	assertSet(t, x, []string{"c"})
}

func TestUnion(t *testing.T) {
	s := readysetgo.Union(readysetgo.New(1, 2, 3), readysetgo.New(3, 4, 5),
		readysetgo.New(5, 6, 7))
	assertSet(t, s, []int{1, 2, 3, 4, 5, 6, 7})
}

func TestIntersection(t *testing.T) {
	s := readysetgo.Intersection(readysetgo.New(1, 2, 3, 4), readysetgo.New(2, 3, 4, 5),
		readysetgo.New(3, 4, 5, 5))
	assertSet(t, s, []int{3, 4})
}

// assertSet checks that the set contains all the expected items and only them
func assertSet[T comparable](t *testing.T, s readysetgo.Set[T], items []T) {
	require.Len(t, s, len(items), "expected length %d", len(items))
	for _, item := range items {
		_, ok := s[item]
		require.True(t, ok, "expected to find %v in set", item)
	}
}
