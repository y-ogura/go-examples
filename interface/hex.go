package main

import "fmt"

// Hex intger type.
type Hex int

// String string func.
func (h Hex) String() string {
	return fmt.Sprintf("%v", int(h))
}
