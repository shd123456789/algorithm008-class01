// 合并2个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    prev := &ListNode{Val:0, Next:nil}
    tmp := prev
    for l1 != nil && l2 != nil {
        if l1.Val > l2.Val {
            tmp.Next = l2
            tmp,l2 = l2,l2.Next 
        } else {
            tmp.Next = l1
            tmp,l1 = l1,l1.Next
        }
    }
    if l1 == nil {
        tmp.Next = l2
    }
    if l2 == nil {
        tmp.Next = l1
    }
    return prev.Next
}