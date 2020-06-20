package main

import "fmt"

type ListNode234 struct {
	Val  int
	Next *ListNode234
}

//判断是否是回文链表
func main() {
	testListNode := &ListNode234{
		Val: 1,
		Next: &ListNode234{
			Val: 2,
			Next: &ListNode234{
				Val: 2,
				Next: &ListNode234{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(isPalindromeList(testListNode))
}

func isPalindromeList(head *ListNode234) bool {
	//一开始以为是数组或字符串，直接想的对撞指针，才发现是链表，那感觉可以先把链表里的值遍历到数组里，再用双指针
	//但是我感觉用栈会更好？
	array := []int{}
	if head == nil || head.Next == nil {
		return true
	}

	for {
		if head == nil {
			break
		}
		array = append(array, head.Val)
		head = head.Next
	}

	//数组的左，右位置
	left := 0
	right := len(array) - 1

	for left <= right {
		if array[left] == array[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}
