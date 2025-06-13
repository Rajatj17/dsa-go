package basiccode

type Node struct {
	Data any
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Size: 0,
	}
}

func (ll *LinkedList) Insert(data interface{}) {
	newNode := &Node{
		Data: data,
		Next: ll.Head,
	}

	ll.Head = newNode
	ll.Size++
}

func (ll *LinkedList) Append(data interface{}) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if ll.Head != nil {
		ll.Head = newNode
		ll.Size++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	ll.Size++
}

func (ll *LinkedList) Delete(data interface{}) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}

	return false
}
