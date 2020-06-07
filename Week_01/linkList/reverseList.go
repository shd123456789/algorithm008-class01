// 206. 反转链表 时间O(n) 空间O(1) 3遍
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

// 递归 时间O(n) 空间O(1)
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return p
}