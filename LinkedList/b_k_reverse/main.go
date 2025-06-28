package kreverse

import basiccode "DSA/LinkedList/a_basic_code"

func KReverse(head *basiccode.Node, k int) *basiccode.Node {
	if head == nil {
		return head
	}

	var newHead *basiccode.Node
	var tail *basiccode.Node

	currentHead := head
	for currentHead != nil {
		var prev *basiccode.Node
		var next *basiccode.Node
		groupHead := currentHead

		count := 0
		for currentHead != nil && count < k {
			next = currentHead.Next
			currentHead.Next = prev
			prev = currentHead
			currentHead = next
		}

		// If newHead is null, set it to the
		// last node of the first group
		if newHead == nil {
			newHead = prev
		}

		// Connect the previous group to the
		// current reversed group
		if tail != nil {
			tail.Next = prev
		}

		// Move tail to the end of the
		// reversed group
		tail = groupHead
	}

	return newHead
}

func KReverseRecursion(head *basiccode.Node, k int) *basiccode.Node {
	if head == nil {
		return head
	}

	currentHead := head
	var prev *basiccode.Node
	var next *basiccode.Node

	count := 0
	for currentHead != nil && count < k {
		next = currentHead.Next
		currentHead.Next = prev
		prev = currentHead
		currentHead = next
	}

	if currentHead != nil {
		head.Next = KReverse(currentHead, k)
	}

	return prev
}
