package main

import "fmt"

type Hex int

func (h Hex) String() {
	return fmt.Sprintf("0x%x", int(h))
}
