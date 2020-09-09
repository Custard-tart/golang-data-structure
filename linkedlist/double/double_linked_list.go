package double

// 单链表：增删改查

// LinkedListNode 双向链表节点
type LinkedListNode struct {
	Pre   *LinkedListNode
	Next  *LinkedListNode
	Value float64
}

// LinkedList 双向链表头节点
type LinkedList struct {
	Header *LinkedListNode
}

// IsEmpty 判断链表是否为空
func (ll *LinkedList) IsEmpty() bool {
	if ll.Header == nil {
		return true
	}
	return false
}

// Length 获取链表的长度
func (ll *LinkedList) Length() int {
	count := 0
	cur := ll.Header
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// Add 从头部添加节点
func (ll *LinkedList) Add(value float64) {
	node := &LinkedListNode{
		Value: value,
		Next:  ll.Header,
	}
	ll.Header = node
}

// Append 从尾部添加节点
func (ll *LinkedList) Append(value float64) {
	if ll.IsEmpty() {
		ll.Header = &LinkedListNode{
			Value: value,
		}
		return
	}
	header := ll.Header
	for header.Next != nil {
		header = header.Next
	}
	header.Next = &LinkedListNode{
		Value: value,
	}
}

// Insert 在指定位置添加节点
func (ll *LinkedList) Insert(index int, value float64) {
	switch {
	// 如果index小于0，则在头部添加
	case index <= 0:
		ll.Add(value)
	// 如果index大于当前链表长度，则在尾部添加
	case index > ll.Length():
		ll.Append(value)
	default:
		count := 0
		var pre *LinkedListNode
		node := ll.Header
		for count < index {
			pre = node
			node = node.Next
			count++
		}
		pre.Next = &LinkedListNode{
			Pre:   pre,
			Next:  node,
			Value: value,
		}
	}
}

// RemoveByIndex 删除指定位置的节点,并返回该节点
func (ll *LinkedList) RemoveByIndex(index int) (res *LinkedListNode) {
	switch {

	// index小于等于0，则删除头部节点
	case index <= 0:
		res = ll.Header
		ll.Header.Next.Pre = nil
		ll.Header = ll.Header.Next

	// index大于当前链表长度，则删除尾部节点
	case index >= ll.Length():
		var pre *LinkedListNode
		node := ll.Header
		for node.Next != nil {
			pre = node
			node = node.Next
		}
		res = node
		pre.Next = nil
	default:
		var pre *LinkedListNode
		node := ll.Header
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
func (ll *LinkedList) RemoveByValue(value float64) (res *LinkedListNode) {
	var pre *LinkedListNode
	node := ll.Header
	for node != nil {
		if node.Value != value {
			pre = node
			node = node.Next
			continue
		}
		if pre == nil {
			ll.Header = node.Next
		} else {
			pre.Next = node.Next
			node.Next.Pre = pre
		}
		res = node
		break
	}
	return res
}
