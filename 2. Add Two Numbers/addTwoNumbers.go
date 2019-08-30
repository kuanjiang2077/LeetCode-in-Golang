/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int
    dummy := &ListNode{}
    p := dummy
    
    for l1 != nil || l2 != nil {
        if p.Next == nil {
            p.Next = &ListNode{}
        }
        p = p.Next
        var n1, n2 int
        if l1 != nil {
            n1 = l1.Val
        }
        if l2 != nil {
            n2 = l2.Val
        }
        sum := n1 + n2 + carry
        p.Val, carry = sum % 10, sum / 10
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    
    if carry > 0 {
        p.Next = &ListNode {
            Val: carry,
        }
    }
    
    return dummy.Next
}