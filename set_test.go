package readysetgo_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ilmaruk/readysetgo"
)

func TestSet_Add(t *testing.T) {
	s := readysetgo.NewSet[int]()

	// Single item
	s.Add(1)
	assertSet(t, s, []int{1})

	// Multiple items
	s.Add(1, 2, 3)
	assertSet(t, s, []int{1, 2, 3})
}

func TestSet_Clear(t *testing.T) {
	s := readysetgo.NewSet(1, 2, 3)
	s.Clear()
	assertSet(t, s, []int{})
}

func TestSet_Copy(t *testing.T) {
	l := readysetgo.NewSet(1, 2, 3)
	r := l.Copy()
	require.NotEqual(t, reflect.ValueOf(l).Pointer() == reflect.ValueOf(r).Pointer(), "expected different sets")
	assertSet(t, r, []int{1, 2, 3})
}

func TestSet_Has(t *testing.T) {
	s := readysetgo.NewSet(1)
	require.True(t, s.Has(1), "expected to find 1 in set")
	require.False(t, s.Has(2), "expected not to find 2 in set")
}

func TestSer_Items(t *testing.T) {
	s := readysetgo.NewSet(1, 2, 3)
	require.ElementsMatch(t, []int{1, 2, 3}, s.Items(), "expected items to match")
}

func TestSet_Update(t *testing.T) {
	s := readysetgo.NewSet(1, 2, 3)
	s.Update(readysetgo.NewSet(3, 4, 5), readysetgo.NewSet(5, 6, 7))
	assertSet(t, s, []int{1, 2, 3, 4, 5, 6, 7})
}

func TestDifference(t *testing.T) {
	s := readysetgo.Difference(readysetgo.NewSet(1, 2, 3, 4), readysetgo.NewSet(3, 4, 5),
		readysetgo.NewSet(1, 3, 6))
	assertSet(t, s, []int{2})
}

func TestSet_IsSubset(t *testing.T) {
	s := readysetgo.NewSet(1, 2, 3)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is subset",
			other:    readysetgo.NewSet(0, 1, 2, 3, 4),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.NewSet(1, 2, 3),
			expected: true,
		},
		{
			name:     "is not subset",
			other:    readysetgo.NewSet(1, 2, 4),
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
	s := readysetgo.NewSet(1, 2, 3)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is disjoint",
			other:    readysetgo.NewSet(0, 4, 5),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.NewSet(1, 2, 3),
			expected: false,
		},
		{
			name:     "is not disjoint",
			other:    readysetgo.NewSet(3, 4, 5),
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
	s := readysetgo.NewSet(0, 1, 2, 3, 4)

	tt := []struct {
		name     string
		other    readysetgo.Set[int]
		expected bool
	}{
		{
			name:     "is superset",
			other:    readysetgo.NewSet(1, 2, 3),
			expected: true,
		},
		{
			name:     "equal sets",
			other:    readysetgo.NewSet(0, 1, 2, 3, 4),
			expected: true,
		},
		{
			name:     "is not superset",
			other:    readysetgo.NewSet(1, 2, 5),
			expected: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, s.IsSuperset(tc.other), "expected %v to be superset of %v", s, tc.other)
		})
	}
}

func TestUnion(t *testing.T) {
	s := readysetgo.Union(readysetgo.NewSet(1, 2, 3), readysetgo.NewSet(3, 4, 5),
		readysetgo.NewSet(5, 6, 7))
	assertSet(t, s, []int{1, 2, 3, 4, 5, 6, 7})
}

func TestIntersection(t *testing.T) {
	s := readysetgo.Intersection(readysetgo.NewSet(1, 2, 3, 4), readysetgo.NewSet(2, 3, 4, 5),
		readysetgo.NewSet(3, 4, 5, 5))
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
