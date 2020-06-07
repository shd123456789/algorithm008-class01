/**
 * 141. 环形链表 hash解法 时间空间复杂度均为O(n) (3遍)
 */
func hasCycle(head *ListNode) bool {
    hash := make(map[*ListNode]int)

    for head != nil {
        if _,ok := hash[head]; ok {
            return true
        }

        hash[head] = 1
        head = head.Next
    }

    return false
}

/**
 * 141. 环形链表 快慢指针 时间复杂度为O(n)空间复杂度为O(1)
 */
func hasCycle(head *ListNode) bool {
    slow,fast := head,head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if fast == slow {
            return true
        }
    }

    return false
}

