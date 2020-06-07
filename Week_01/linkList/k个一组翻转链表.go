// 时间复杂度为O(n) 空间复杂度为O(1) 1遍
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k <= 1 {
        return head
    }
    res := &ListNode{Next:head}
    idx,prev,curr := 1,res,head
    for curr != nil {
        if idx % k == 0 {
            node := reverse(prev.Next, curr.Next)
            tmp := prev.Next
            prev.Next = node
            prev = tmp
            curr = prev.Next
        } else {
            curr = curr.Next
        }
        idx++
    }
    return res.Next
}
func reverse(head *ListNode, end *ListNode) *ListNode {
    prev := end
    for head != end {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}