package main

import (
	"fmt"
	"github.com/sonntuet1997/patricia-tree/src"
)

func main() {
	t := src.NewNodeFromHex("cc")
	x := src.NewNodeFromHex("11")
	fmt.Println(t.Hash())
	fmt.Println(x.Hash())
}
