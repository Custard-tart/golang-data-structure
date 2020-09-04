package linkedlist

import (
	"errors"
)

// 单链表：增删改查

// SingleChainListNode 单链表节点
type SingleChainListNode struct {
	Next  *SingleChainListNode
	Value float64
}

// SingleChainList 单链表头节点
type SingleChainList struct {
	Header *SingleChainListNode
}

// IsEmpty 判断链表是否为空
func (scl *SingleChainList) IsEmpty() bool {
	if scl.Header == nil {
		return true
	}
	return false
}

// Length 获取链表的长度
func (scl *SingleChainList) Length() int {
	count := 0
	cur := scl.Header
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 从头部添加节点
func (scl *SingleChainList) Add(value float64) {
	node := &SingleChainListNode{
		Value: value,
		Next:  scl.Header,
	}
	scl.Header = node
}

// Append 从尾部添加节点
func (scl *SingleChainList) Append(value float64) {
	if scl.IsEmpty() {
		scl.Header = &SingleChainListNode{
			Value: value,
		}
		return
	}
	header := scl.Header
	for header.Next != nil {
		header = header.Next
	}
	header.Next = &SingleChainListNode{
		Value: value,
	}
}

// Insert 在指定位置添加节点
func (scl *SingleChainList) Insert(index int, value float64) {
	switch {
	// 如果index小于0，则在头部添加
	case index <= 0:
		scl.Add(value)
	// 如果index大于当前链表长度，则在尾部添加
	case index > scl.Length():
		scl.Append(value)
	default:
		count := 0
		pre := scl.Header
		for count < index {
			pre = pre.Next
			count++
		}
		pre.Next = &SingleChainListNode{
			Next:  pre.Next,
			Value: value,
		}
	}
}

// RemoveByIndex 删除指定位置的节点,并返回该节点
func (scl *SingleChainList) RemoveByIndex(index int) (res *SingleChainListNode) {
	switch {
	// index小于等于0，则删除头部节点
	case index <= 0:
		res = scl.Header
		scl.Header = scl.Header.Next
	// index大于当前链表长度，则删除尾部节点
	case index >= scl.Length():
		var pre *SingleChainListNode
		node := scl.Header
		for node.Next != nil {
			pre = node
			node = node.Next
		}
		res = node
		pre.Next = nil
	default:
		var pre *SingleChainListNode
		node := scl.Header
		count := 0
		for count < index {
			pre = node
			node = node.Next
			count++
		}
		res = node
		pre.Next = node.Next
	}
	return res
}

// RemoveByValue 删除指定值的节点
func (scl *SingleChainList) RemoveByValue(value float64) (res *SingleChainListNode) {
	var pre *SingleChainListNode
	node := scl.Header
	for node != nil {
		if node.Value != value {
			pre = node
			node = node.Next
			continue
		}
		if pre == nil {
			scl.Header = node.Next
		} else {
			pre.Next = node.Next
		}
		res = node
		break
	}
	return res
}

// GetNodeByIndex 根据index获取node
func (scl *SingleChainList) GetNodeByIndex(index int) (res *SingleChainListNode, err error) {
	count := 0
	node := scl.Header
	if scl.Length() <= index {
		return res, errors.New("Out of index")
	}
	if index < 0 {
		return res, errors.New("Index must be greater than or equal to 0")
	}
	for count < index {
		node = node.Next
	}
	res = node
	return res, err
}
