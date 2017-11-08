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

//InsertEnd adds element to the end of list
func (s *Sll) InsertEnd(v interface{}) {
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

//InsertFront adds element to the front of list
func (s *Sll) InsertFront(v interface{}) {
	n := &node{Value: v}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n

	}

	s.length++
}

//InsertAfter adds element after specified element
func (s *Sll) InsertAfter(v interface{}, after interface{}) error {
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

//RemoveEnd removes element from the end of list
func (s *Sll) RemoveEnd() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}
	temp1 := s.head
	var temp2 *node
	if temp1.next == nil {
		s.head = nil
		s.length--
		return nil
	}

	for temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}

	temp2.next = nil
	s.length--
	return nil
}

//RemoveFront adds element to the end of list
func (s *Sll) RemoveFront() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}

	temp := s.head

	if temp.next == nil {
		s.head = nil
		s.length--
		return nil
	}

	s.head = temp.next
	s.length--
	return nil
}

//RemoveSpecific removes specified element from the list
func (s *Sll) RemoveSpecific(v interface{}) error {
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

//Display displays the list
func (s *Sll) Display() {
	if s.head == nil {
		fmt.Print("List is empty.")
	}
	cur := s.head
	for cur != nil {
		fmt.Printf("%v ", cur.Value)
		cur = cur.next
	}
}

//Size returns length of the list
func (s *Sll) Size() int {
	return s.length
}
