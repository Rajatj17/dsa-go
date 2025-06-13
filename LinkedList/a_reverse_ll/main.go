package reversell

import (
	basiccode "DSA/LinkedList/a_basic_code"
)

func RecReverse(head *basiccode.Node) *basiccode.Node {
	if head == nil || head.Next == nil {
		return head
	}

	startHead := RecReverse(head.Next)

	head.Next.Next = head
	head.Next = nil

	return startHead
}

func IterativeReverse(head *basiccode.Node) *basiccode.Node {
	var prev *basiccode.Node
	for head.Next != nil {
		nextHead := head.Next
		head.Next = prev
		prev = head
		head = nextHead
	}

	return head
}
