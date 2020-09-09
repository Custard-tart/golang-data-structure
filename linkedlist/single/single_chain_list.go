package single

import (
	"errors"
)

// 单链表：增删改查

// ChainListNode 单链表节点
type ChainListNode struct {
	Next  *ChainListNode
	Value float64
}

// ChainList 单链表头节点
type ChainList struct {
	Header *ChainListNode
}

// IsEmpty 判断链表是否为空
func (cl *ChainList) IsEmpty() bool {
	if cl.Header == nil {
		return true
	}
	return false
}

// Length 获取链表的长度
func (cl *ChainList) Length() int {
	count := 0
	cur := cl.Header
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 从头部添加节点
func (cl *ChainList) Add(value float64) {
	node := &ChainListNode{
		Value: value,
		Next:  cl.Header,
	}
	cl.Header = node
}

// Append 从尾部添加节点
func (cl *ChainList) Append(value float64) {
	if cl.IsEmpty() {
		cl.Header = &ChainListNode{
			Value: value,
		}
		return
	}
	header := cl.Header
	for header.Next != nil {
		header = header.Next
	}
	header.Next = &ChainListNode{
		Value: value,
	}
}

// Insert 在指定位置添加节点
func (cl *ChainList) Insert(index int, value float64) {
	switch {
	// 如果index小于0，则在头部添加
	case index <= 0:
		cl.Add(value)
	// 如果index大于当前链表长度，则在尾部添加
	case index > cl.Length():
		cl.Append(value)
	default:
		count := 0
		pre := cl.Header
		for count < index {
			pre = pre.Next
			count++
		}
		pre.Next = &ChainListNode{
			Next:  pre.Next,
			Value: value,
		}
	}
}

// RemoveByIndex 删除指定位置的节点,并返回该节点
func (cl *ChainList) RemoveByIndex(index int) (res *ChainListNode) {
	switch {
	// index小于等于0，则删除头部节点
	case index <= 0:
		res = cl.Header
		cl.Header = cl.Header.Next
	// index大于当前链表长度，则删除尾部节点
	case index >= cl.Length():
		var pre *ChainListNode
		node := cl.Header
		for node.Next != nil {
			pre = node
			node = node.Next
		}
		res = node
		pre.Next = nil
	default:
		var pre *ChainListNode
		node := cl.Header
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
func (cl *ChainList) RemoveByValue(value float64) (res *ChainListNode) {
	var pre *ChainListNode
	node := cl.Header
	for node != nil {
		if node.Value != value {
			pre = node
			node = node.Next
			continue
		}
		if pre == nil {
			cl.Header = node.Next
		} else {
			pre.Next = node.Next
		}
		res = node
		break
	}
	return res
}

// GetNodeByIndex 根据index获取node
func (cl *ChainList) GetNodeByIndex(index int) (res *ChainListNode, err error) {
	count := 0
	node := cl.Header
	if cl.Length() <= index {
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
