/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    var res []int
    for  head != nil {
        res = append(res, head.Val)
        head = head.Next
    }
    i,j := 0,len(res) - 1
    for i < j {
        res[i], res[j] = res[j],res[i]
        i++
        j--
    }
    return res
}