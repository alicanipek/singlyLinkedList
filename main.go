package main

import "fmt"

type node struct {
	Value interface{}

	next *node
}

type sll struct {
	head *node

	length int
}

func newSll(h *node, length int) *sll {
	var s sll
	s.head = h
	s.length = length
	return &s
}

func (s *sll) insertEnd(v interface{}) {
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

func (s *sll) insertStart(v interface{}) {
	n := &node{Value: v}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n

	}

	s.length++
}

func (s *sll) insertAfter(v interface{}, after interface{}) error {
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

func (s *sll) removeEnd() (*node, error) {
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

func (s *sll) removeFront() (*node, error) {
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

func (s *sll) removeSpecific(v interface{}) error {
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

func (s *sll) display() {
	if s.head == nil {
		fmt.Print("List is empty.")
	}
	cur := s.head
	for cur != nil {
		fmt.Printf("%v ", cur.Value)
		cur = cur.next
	}
}

func main() {
	s := newSll(nil, 0)

	s.insertEnd(1)
	s.insertEnd(2)
	s.insertEnd(3)
	s.insertEnd(4)
	s.insertEnd(5)

	s.display()

	/*for s.length != 0 {
		e, er := s.removeEnd()
		if er != nil {
			fmt.Printf("error")
		}
		fmt.Println(e)
	}*/

	/*s.removeFront()
	s.display()
	*/

	s.removeSpecific(4)
	s.display()

	s.insertAfter(8, 2)
	s.display()
}
