package main

import (
	"fmt"
	"how2useinterfaces/animal"
	"how2useinterfaces/circus"
)

// 1. Define the types
// 2. Define the interface at point of use.
func main() {

	dog := animal.Dog{}
	a   := circus.Perform(dog)

	// compile error
	// cat := animal.Cat{}
	// b   := circus.Perform(cat) // cannot use animal.Cat literal (type animal.Cat) as type circus.Speaker in argument to circus.Perform: animal.Cat does not implement circus.Speaker (missing Speaks method)
	fmt.Println(a)
}
