// 24. 两两交换链表中的节点 时间复杂度O(n) 空间复杂度O(1) 2遍
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