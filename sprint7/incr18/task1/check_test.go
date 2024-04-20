package check

import (
	"fmt"
)

func Example() {
	var s Student
	s.SetName("prosto")
	fmt.Println(s.GetName())

	// Output:
	// Prosto
}
