package main

import "fmt"

type LinkedList[T any] interface {
	Length() int
	InsertAt(item *T, index int)
	Remove(item *T) *T
	RemoveAt(index int) *T
	Append(item *T)
	Prepend(item *T)
	Get(index int) *T
}

type Node struct {
	Value int
	Next  *Node
}

type IntLinkedList struct {
	head *Node
	tail *Node
}

func (l *IntLinkedList) Length() int {
	length := 0
	head := l.head
	for head != nil {
		length++
		head = head.Next
	}

	return length
}

func (l *IntLinkedList) InsertAt(item *Node, index int) {
	cur := l.head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	if cur.Next == nil {
		l.tail = item
	}

	cur.Next = item
}

func (l *IntLinkedList) Remove(item *Node) *Node {
	if l.head == nil {
		return nil
	}

	if l.head == item {
		res := l.head
		l.head = l.head.Next

		if l.head.Next == nil {
			l.tail = nil
		}

		return res
	}

	cur := l.head

	for cur.Next != nil && cur.Next.Value != item.Value {
		cur = cur.Next
	}

	res := cur.Next

	cur.Next = cur.Next.Next

	if cur.Next == nil {
		l.tail = cur
	}

	return res
}

func (l *IntLinkedList) RemoveAt(index int) *Node {
  if l.head == nil {
    return nil
  }

  if index == 0 {
    res := l.head
    l.head = l.head.Next 

    if l.head.Next == nil {
      l.tail = nil
    }

    return res
  }

	cur := l.head

	for i := 0; i < index -1; i++ {
		cur = cur.Next
	}

	res := cur
	cur.Next = cur.Next.Next

  if cur.Next == nil {
    l.tail = nil
  }

	return res
}

func (l *IntLinkedList) Append(item *Node) *Node {
	if l.head == nil {
		l.head = item
		l.tail = item
		return l.tail
	}

	cur := l.tail

	cur.Next = item

	l.tail = item

	return l.tail
}

func (l *IntLinkedList) Prepend(item *Node) *Node {
	if l.head == nil {
		l.head = item
		l.tail = item
		return l.head
	}

	item.Next = l.head

	l.head = item

	return l.head
}

func (l *IntLinkedList) Get(index int) *Node {
	cur := l.head

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur
}

func main() {
	ill := new(IntLinkedList)

	ill.Append(&Node{Value: 10})
	ill.Append(&Node{Value: 20})
	ill.Append(&Node{Value: 30})

	fmt.Println(ill.Length())

	fmt.Println(ill.Get(0))
	fmt.Println(ill.Get(1))
	fmt.Println(ill.Get(2))

	ill.Remove(&Node{Value: 20})

	fmt.Println(ill.Length())

	fmt.Println(ill.Get(0))
	fmt.Println(ill.Get(1))

	ill.Prepend(&Node{Value: 7})

	fmt.Println(ill.Length())

	fmt.Println(ill.Get(0))
	fmt.Println(ill.Get(1))
	fmt.Println(ill.Get(2))

	ill.RemoveAt(0)

	fmt.Println(ill.Length())
	fmt.Println(ill.Get(0))
	fmt.Println(ill.Get(1))

	ill.RemoveAt(1)

	fmt.Println(ill.Length())
	fmt.Println(ill.Get(0))
	fmt.Println(ill.Get(1))

}
