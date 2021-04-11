/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    iter := &ListNode{}
    answer := iter
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            iter.Next = l1
            l1 = l1.Next
        } else {
            iter.Next = l2
            l2 = l2.Next
        }
        iter = iter.Next
    }
    for l1 != nil {
        iter.Next = l1
        l1 = l1.Next
        iter = iter.Next
    }
    for l2 != nil {
        iter.Next = l2
        l2 = l2.Next
        iter = iter.Next
    }
    return answer.Next
}
