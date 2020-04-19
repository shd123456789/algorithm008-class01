// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
    pre := &ListNode{Val:0, Next:head}
    tmp := pre

    for tmp.Next != nil && head.Next != nil {
        firstNode := tmp.Next
        secondNode := head.Next

        firstNode.Next = secondNode.Next
        secondNode.Next = firstNode
        tmp.Next = secondNode

        tmp = firstNode
        head = firstNode.Next
    }

    return pre.Next
}