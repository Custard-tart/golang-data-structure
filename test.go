package main

import (
	"fmt"

	"github.com/Custard-tart/golang-data-structure/linkedlist/double"
	"github.com/Custard-tart/golang-data-structure/linkedlist/single"
)

func main() {
	// SingleChainListTest()
	// DoubleLinkedListTest()
}

// SingleChainListTest 测试单链表
func SingleChainListTest() {
	a := single.ChainList{}
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)
	a.Append(5)
	a.Append(6)

	a.Add(0)
	a.Append(7)
	a.Insert(4, 4.5)
	a.RemoveByIndex(11)
	a.RemoveByValue(1)

	node := a.Header
	for node != nil {
		fmt.Println("node--------------", node.Value)
		node = node.Next
	}
}

// DoubleLinkedListTest 测试双向链表
func DoubleLinkedListTest() {
	a := double.LinkedList{}
	a.Append(1)
	a.Append(2)
	a.Append(3)
	a.Append(4)
	a.Append(5)
	a.Append(6)

	a.Add(0)
	a.Append(7)
	a.Insert(3, 3.5)

	node := a.Header
	for node != nil {
		fmt.Println("node--------------", node.Value)
		node = node.Next
	}
}
