/**
 *
 * File:    leetcode2.go
 *          两数相加
 * 
 * Author:  yiluqingtai(1572236483@qq.com)
 *          Created on 21/7/8
 *          
 **/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head *ListNode
    var tail *ListNode
    carry := 0

    for l1 != nil || l2 != nil {
        n1, n2 := 0, 0
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }
        sum := n1 + n2 + carry
        sum, carry = sum % 10, sum / 10
        if head == nil {
            head = &ListNode{Val : sum}
            tail = head
        } else {
            tail.Next = &ListNode{Val : sum}
            tail = tail.Next
        }
    }

    if carry > 0 {
        tail.Next = &ListNode{Val : carry}
    }
    return head
}