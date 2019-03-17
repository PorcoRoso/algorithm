package linkedList

import (
	"fmt"

)

type ListNode struct {
	Val int
	Next *ListNode
}

// MergeTwoLists leetcode 21.
// 哨兵
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2

	p := &ListNode{
		-1,
		nil,
	}

	q := p

	for ;p1 != nil && p2 != nil; {

		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	for ;p1 != nil; {
		p.Next = p1
		p1 = p1.Next
		p = p.Next
	}

	for ;p2 != nil; {
		p.Next = p2
		p2 = p2.Next
	}
	return q.Next
}

// RemoveNthFromEnd leetcode 19.
// 哨兵
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}
	cur := head
	m := 0
	for ;cur != nil; {
		cur = cur.Next
		m++
	}

	if m < n {
		return head
	}

	var p = & ListNode {
		-1,
		head,
	}


	cur = head
	pre := p

	for i := 0; i < m-n && cur != nil; i++ {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = cur.Next
	return p.Next
}

// HasCycle leetcode 141
// P1 step1, P2 step 2, finally p2 run after p1
func hasCycle(head *ListNode) bool {

	p1 := head
	p2 := p1
	for ;p1 != nil && p2 != nil; {
		p1 = p1.Next
		p2 = p2.Next
		if p2 != nil {
			p2 = p2.Next
		} else {
			break
		}
		if p1 == p2 {
			return true
		}
	}
	return false
}

// ReverseList leetcode 206
// pre, cur, aft
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre = &ListNode {}

	pre = nil
	cur := head
	aft := cur.Next

	for ;cur.Next!= nil; {
		aft = cur.Next
		cur.Next = pre
		pre = cur
		cur = aft
	}
	cur.Next = pre
	return cur
}

// MiddleNode leetcode 876
func middleNode(head *ListNode) *ListNode {

	n := 0
	for p:=head; p!= nil; {
		p = p.Next
		n++
	}

	n = n/2

	for ;n>0; {
		head = head.Next
		n--
	}

	return head
}

//ShowList for debug
func showList(list *ListNode) {
	if list == nil {
		fmt.Printf("ListNode is nil")
	}

	for ;list != nil; {
		fmt.Printf("> %d ", list.Val)
		list = list.Next
	}
	fmt.Printf(" > nil ")
}

func main() {

	p3 := &ListNode{
		9,
		nil,
	}


	p2 := &ListNode{
		7,
		p3,
	}
	p1 := &ListNode{
		5,
		p2,
	}

	//r := mergeTwoLists(p1, l1)
	r1 := reverseList(p1)
	showList(r1)

}