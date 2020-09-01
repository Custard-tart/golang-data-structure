package linkedlist

import (
	"errors"
)

// 单链表：增删改查

// SingleChainListNode 单链表节点
type SingleChainListNode struct {
	Next  *SingleChainListNode
	Value int
}

// SingleChainListHeader 但链表头节点
type SingleChainListHeader struct {
	Header *SingleChainListNode
}

// IsEmpty 判断链表是否为空
func (sclh *SingleChainListHeader) IsEmpty() bool {
	if sclh.Header == nil {
		return true
	}
	return false
}

// Length 获取链表的长度
func (sclh *SingleChainListHeader) Length() int {
	count := 0
	cur := sclh.Header
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 从头部添加节点
func (sclh *SingleChainListHeader) Add(value int) {
	node := &SingleChainListNode{
		Value: value,
		Next:  sclh.Header,
	}
	sclh.Header = node
}

// Append 从尾部添加节点
func (sclh *SingleChainListHeader) Append(value int) {
	if sclh.IsEmpty() {
		sclh.Header = &SingleChainListNode{
			Value: value,
		}
		return
	}
	header := sclh.Header
	for header.Next != nil {
		header = header.Next
	}
	header.Next = &SingleChainListNode{
		Value: value,
	}
}

// Insert 在指定位置添加节点
func (sclh *SingleChainListHeader) Insert(index int, value int) {
	switch {
	// 如果index小于0，则在头部添加
	case index <= 0:
		sclh.Add(value)
	// 如果index大于当前链表长度，则在尾部添加
	case index > sclh.Length():
		sclh.Append(value)
	default:
		count := 0
		pre := sclh.Header
		for count < index-1 {
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
func (sclh *SingleChainListHeader) RemoveByIndex(index int) (res *SingleChainListNode) {
	switch {
	// index小于等于0，则删除头部节点
	case index <= 0:
		res = sclh.Header
		sclh.Header = sclh.Header.Next
	// index大于当前链表长度，则删除尾部节点
	case index >= sclh.Length():
		var pre *SingleChainListNode
		node := sclh.Header
		for node != nil {
			pre = node
			node = node.Next
		}
		res = node
		pre.Next = nil
	default:
		var pre *SingleChainListNode
		node := sclh.Header
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
func (sclh *SingleChainListHeader) RemoveByValue(value int) (res *SingleChainListNode) {
	var pre *SingleChainListNode
	node := sclh.Header
	for node != nil {
		if node.Value != value {
			pre = node
			node = node.Next
			continue
		}
		if pre == nil {
			sclh.Header = node.Next
		} else {
			pre.Next = node.Next
		}
		res = node
		break
	}
	return res
}

// GetNodeByIndex 根据index获取node
func (sclh *SingleChainListHeader) GetNodeByIndex(index int) (res *SingleChainListNode, err error) {
	count := 0
	node := sclh.Header
	if sclh.Length() <= index {
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
