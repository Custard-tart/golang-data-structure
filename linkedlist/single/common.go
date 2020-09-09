package single

// 该文件实现了一些关于单链表常见的算法

// Reversal 反转链表
func (scl *ChainList) Reversal() {
	var pre *ChainListNode
	node := scl.Header
	for node != nil {
		var next *ChainListNode
		if node.Next != nil {
			next = node.Next
		}
		node.Next = pre
		pre = node
		node = next
	}
	scl.Header = pre
}
