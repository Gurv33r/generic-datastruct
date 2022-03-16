package main

import (
	"fmt"

	"github.com/Gurv33r/generic-datastruct/datastruct"
)

func main() {
	set := datastruct.NewSetFrom[uint](1, 1, 3)
	fmt.Println(set)
}
