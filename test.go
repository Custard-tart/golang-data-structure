package main

import (
	"fmt"

	"github.com/Custard-tart/golang-data-structure/linkedlist"
)

func main() {
	a := linkedlist.SingleChainListHeader{}
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)
	a.Append(5)
	a.Append(6)

	a.Add(0)
	a.Append(7)
	a.Insert(4, 10)
	a.RemoveByIndex(7)
	a.RemoveByValue(1)

	node := a.Header
	for node != nil {
		fmt.Println("node--------------", node.Value)
		node = node.Next
	}
}
