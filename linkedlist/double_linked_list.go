package linkedlist

// 单链表：增删改查

// DoubleLinkedListNode 双向链表节点
type DoubleLinkedListNode struct {
	Pre   *DoubleLinkedListNode
	Next  *DoubleLinkedListNode
	Value float64
}

// DoubleLinkedList 双向链表头节点
type DoubleLinkedList struct {
	Header *DoubleLinkedListNode
}

// IsEmpty 判断链表是否为空
func (dll *DoubleLinkedList) IsEmpty() bool {
	if dll.Header == nil {
		return true
	}
	return false
}

// Length 获取链表的长度
func (dll *DoubleLinkedList) Length() int {
	count := 0
	cur := dll.Header
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 从头部添加节点
func (dll *DoubleLinkedList) Add(value float64) {
	node := &DoubleLinkedListNode{
		Value: value,
		Next:  dll.Header,
	}
	dll.Header = node
}

// Append 从尾部添加节点
func (dll *DoubleLinkedList) Append(value float64) {
	if dll.IsEmpty() {
		dll.Header = &DoubleLinkedListNode{
			Value: value,
		}
		return
	}
	header := dll.Header
	for header.Next != nil {
		header = header.Next
	}
	header.Next = &DoubleLinkedListNode{
		Value: value,
	}
}

// Insert 在指定位置添加节点
func (dll *DoubleLinkedList) Insert(index int, value float64) {
	switch {
	// 如果index小于0，则在头部添加
	case index <= 0:
		dll.Add(value)
	// 如果index大于当前链表长度，则在尾部添加
	case index > dll.Length():
		dll.Append(value)
	default:
		count := 0
		var pre *DoubleLinkedListNode
		node := dll.Header
		for count < index {
			pre = node
			node = node.Next
			count++
		}
		pre.Next = &DoubleLinkedListNode{
			Pre:   pre,
			Next:  node,
			Value: value,
		}
	}
}

// RemoveByIndex 删除指定位置的节点,并返回该节点
func (dll *DoubleLinkedList) RemoveByIndex(index int) (res *DoubleLinkedListNode) {
	switch {

	// index小于等于0，则删除头部节点
	case index <= 0:
		res = dll.Header
		dll.Header.Next.Pre = nil
		dll.Header = dll.Header.Next

	// index大于当前链表长度，则删除尾部节点
	case index >= dll.Length():
		var pre *DoubleLinkedListNode
		node := dll.Header
		for node.Next != nil {
			pre = node
			node = node.Next
		}
		res = node
		pre.Next = nil
	default:
		var pre *DoubleLinkedListNode
		node := dll.Header
		count := 0
		for count < index {
			pre = node
			node = node.Next
			count++
		}
		res = node
		pre.Next = node
		node.Pre = pre
	}
	return res
}

// RemoveByValue 删除指定值的节点
func (dll *DoubleLinkedList) RemoveByValue(value float64) (res *DoubleLinkedListNode) {
	var pre *DoubleLinkedListNode
	node := dll.Header
	for node != nil {
		if node.Value != value {
			pre = node
			node = node.Next
			continue
		}
		if pre == nil {
			dll.Header = node.Next
		} else {
			pre.Next = node.Next
			node.Next.Pre = pre
		}
		res = node
		break
	}
	return res
}
