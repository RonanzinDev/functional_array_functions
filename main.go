package main

import (
	"fmt"

	array "github.com/ronanzindev/functional_array_functions/types"
)

type Person struct {
	name string
}

var porra []Person

func main() {
	porra = []Person{{name: "Gabriel"}, {name: "Ronan"}}
	n := []int{1, 2, 3, 4}
	a := array.NewArray(porra)
	b := a.Filter(func(p Person) bool {
		return p.name == "Gabriel"
	})
	f := array.NewArray(n)
	h := f.Filter(func(i int) bool {
		return i > 2
	})
	fmt.Println(b)
	fmt.Println(h)

}
