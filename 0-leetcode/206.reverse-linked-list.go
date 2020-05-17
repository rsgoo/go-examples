package main

import "fmt"

/**
 * 206. 反转链表
 * 反转一个单链表。
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	testListNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	list := ReverseList(testListNode)
	fmt.Println(list.Val)
	fmt.Println(list.Next.Val)
	fmt.Println(list.Next.Next.Val)
	fmt.Println(list.Next.Next.Next.Val)

}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
