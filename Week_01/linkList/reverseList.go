// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
    var newHead *ListNode

    for head != nil {
        next := head.Next
        head.Next = newHead
        newHead = head
        head = next 
    }

    return newHead
}