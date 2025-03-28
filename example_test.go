package readysetgo_test

import (
	"fmt"

	"github.com/ilmaruk/readysetgo"
)

func ExampleNew() {
	ints := readysetgo.New(1, 2, 3, 2, 1)
	fmt.Println("ints:", ints)
	strs := readysetgo.New("a", "b", "c", "c", "C")
	fmt.Println("strs:", strs)
	// Output:
	// ints: map[1:{} 2:{} 3:{}]
	// strs: map[C:{} a:{} b:{} c:{}]
}

func ExampleSet_Remove() {
	s := readysetgo.New(1, 2, 3, 2, 1)
	fmt.Println("set:", s)

	var res bool
	res = s.Remove(2)
	fmt.Println("remove 2:", res)
	fmt.Println("set:", s)

	res = s.Remove(2)
	fmt.Println("remove 2:", res)
	fmt.Println("set:", s)
	// Output:
	// set: map[1:{} 2:{} 3:{}]
	// remove 2: true
	// set: map[1:{} 3:{}]
	// remove 2: false
	// set: map[1:{} 3:{}]
}

func ExampleSet_DifferenceUpdate() {
	a := readysetgo.New("apple", "banana", "cherry")
	b := readysetgo.New("google", "microsoft", "apple")
	c := readysetgo.New("cherry", "micra", "bluebird")
	a.DifferenceUpdate(b, c)
	fmt.Println("set:", a)
	// Output:
	// set: map[banana:{}]
}

func ExampleSet_IntersectionUpdate() {
	x := readysetgo.New("a", "b", "c")
	y := readysetgo.New("c", "d", "e")
	z := readysetgo.New("f", "g", "c")
	x.IntersectionUpdate(y, z)
	fmt.Println("set:", x)
	// Output:
	// set: map[c:{}]
}
