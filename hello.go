package main

import (
	"fmt"

	"github.com/rohopt/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(morestrings.ReverseRunes("Hello, Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
