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
