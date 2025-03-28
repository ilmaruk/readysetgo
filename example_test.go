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
