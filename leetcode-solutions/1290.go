/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    var answer int
    for head != nil {
        answer = (answer << 1) | head.Val
        head = head.Next
    }
    return answer
}
