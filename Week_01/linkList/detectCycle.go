// 环形链表 II 寻找第环节点
func detectCycle(head *ListNode) *ListNode {
    fast,slow := head, head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if (slow == fast) {
            for head != slow {
                head = head.Next
                slow = slow.Next  
            }
            return head;
        }
    }
    return nil
}