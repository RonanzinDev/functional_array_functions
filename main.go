package main

import (
	"fmt"

	"github.com/ronanzindev/functional_array_functions/types"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := types.NewArray(a)
	b.Append(5)
	fmt.Println(b)
}
