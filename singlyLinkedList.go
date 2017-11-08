package singlyLinkedList

import "fmt"

type node struct {
	Value interface{}

	next *node
}

//Sll singly linked list type
type Sll struct {
	head *node

	length int
}

//NewSll creates a singly linked list
func NewSll() *Sll {
	var s Sll
	s.head = nil
	s.length = 0
	return &s
}

func (s *Sll) insertEnd(v interface{}) {
	n := &node{Value: v}
	if s.head == nil {
		s.head = n
	} else {
		current := s.head

		for current.next != nil {
			current = current.next
		}

		current.next = n
	}
	s.length++
}

func (s *Sll) insertStart(v interface{}) {
	n := &node{Value: v}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n

	}

	s.length++
}

func (s *Sll) insertAfter(v interface{}, after interface{}) error {
	in := &node{Value: v}
	if s.head == nil {
		s.head = in
		return nil
	}

	cur := s.head
	for cur.Value != after {
		cur = cur.next
	}

	if cur.Value == after {
		in.next = cur.next
		cur.next = in
		s.length++
		return nil
	}

	return fmt.Errorf("Cant find after node")
}

func (s *Sll) removeEnd() (*node, error) {
	if s.head == nil {
		return nil, fmt.Errorf("list is empty")
	}
	temp1 := s.head
	var temp2 *node
	if temp1.next == nil {
		s.head = nil
		s.length--
		return temp1, nil
	}

	for temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}

	temp2.next = nil
	s.length--
	return temp1, nil
}

func (s *Sll) removeFront() (*node, error) {
	if s.head == nil {
		return nil, fmt.Errorf("list is empty")
	}

	temp := s.head

	if temp.next == nil {
		s.head = nil
		s.length--
		return temp, nil
	}

	s.head = temp.next
	s.length--
	return temp, nil
}

func (s *Sll) removeSpecific(v interface{}) error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	temp1 := s.head
	var temp2 *node

	for temp1.Value != v {
		if temp1 == nil {
			return fmt.Errorf("given node not found in the listde")
		}
		temp2 = temp1
		temp1 = temp1.next
	}

	temp2.next = temp1.next
	s.length--
	return nil
}

func (s *Sll) display() {
	if s.head == nil {
		fmt.Print("List is empty.")
	}
	cur := s.head
	for cur != nil {
		fmt.Printf("%v ", cur.Value)
		cur = cur.next
	}
}

func (s *Sll) size() int {
	return s.length
}
