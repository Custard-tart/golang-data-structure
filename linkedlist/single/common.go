package single

// 该文件实现了一些关于单链表常见的算法

// ReversalSingleChainList 反转链表
func ReversalSingleChainList(scl *ChainList) {
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

// GetFirstCommonNode 输入两个链表，找出它们的第一个公共结点
func GetFirstCommonNode(cl1, cl2 *ChainList) *ChainListNode {
	node1, node2 := cl1.Header, cl2.Header
	if node1 == nil || node2 == nil {
		return nil
	}
	for node1 != node2 {
		if node1 == nil {
			node1 = cl2.Header
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = cl1.Header
		} else {
			node2 = node2.Next
		}
	}
	return node1

}
