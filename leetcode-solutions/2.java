/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode answer = new ListNode(0);
        ListNode t = answer;
        int carry = 0;
        while (l1 != null || l2 != null) {
            int x = carry;
            if (l1 != null) {
                x += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                x += l2.val;
                l2 = l2.next;
            }
            carry = x / 10;
            t.next = new ListNode(x % 10);
            t = t.next;
        }
        if (carry > 0) {
            t.next = new ListNode(1);
        }
        return answer.next;
    }
}
