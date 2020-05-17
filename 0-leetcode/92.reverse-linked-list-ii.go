package main

import "fmt"

type ListNode92 struct {
	Val  int
	Next *ListNode92
}

func main() {
	testListNode := &ListNode92{
		Val: 1,
		Next: &ListNode92{
			Val: 2,
			Next: &ListNode92{
				Val: 3,
				Next: &ListNode92{
					Val: 4,
					Next: &ListNode92{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	node := (reverseBetween(testListNode, 1, 3))
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func reverseBetween(head *ListNode92, m int, n int) *ListNode92 {
	arr := []*ListNode92{}

	if head == nil {
		return nil
	}

	for {
		if head == nil {
			break
		}
		arr = append(arr, head)
		head = head.Next
	}
	//fmt.Println(arr)

	arrMN := arr[m : n+1]

	arrMNLen := len(arrMN)
	for i := 0; i < arrMNLen/2; i++ {
		arrMN[i], arrMN[arrMNLen-1] = arrMN[arrMNLen-1], arrMN[i]
	}
	arr1 := []*ListNode92{}
	arr1 = append(arr1, arr[:m]...)
	arr1 = append(arr1, arrMN...)
	arr1 = append(arr1, arr[n+1:]...)
	//fmt.Println(arr1)
	//fmt.Println("--------------")

	//将元素为链表的数组转换为链表
	root := arr1[0]
	other := root
	for i := 1; i < len(arr1); i++ {
		tempNode := arr1[i]
		other.Next = tempNode
		other = tempNode
	}
	return root
}
