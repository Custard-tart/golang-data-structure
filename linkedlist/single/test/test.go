package main

import (
	"fmt"

	"github.com/Custard-tart/golang-data-structure/linkedlist/single"
)

func main() {
	a := single.ChainList{}
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)
	a.Append(5)
	a.Append(6)

	single.ReversalSingleChainList(&a)

	node := a.Header
	for node != nil {
		fmt.Println("node--------------", node.Value)
		node = node.Next
	}
}
