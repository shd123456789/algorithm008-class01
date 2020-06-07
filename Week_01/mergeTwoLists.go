// 合并2个有序链表(第2遍) 时间复杂度为O(n) 空间为O(1) 2遍
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    prev := &ListNode{}
    curr := prev
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            curr.Next = l1
            curr = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            curr = l2
            l2 = l2.Next
        }
    }
    if l1 != nil {
        curr.Next = l1
    }
    if l2 != nil {
        curr.Next = l2
    }
    return prev.Next
}