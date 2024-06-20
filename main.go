package main

import "fmt"

type Stack struct {
	head *node
}

type node struct {
	val  string
	next *node
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(el string) {
	s.head = &node{
		val:  el,
		next: s.head,
	}
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.head = s.head.next
}

func (s *Stack) Head() string {
	if s.IsEmpty() {
		return ""
	}
	return s.head.val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

// Main

func main() {
	st := New()

	st.Push("a")
	st.Push("b")
	fmt.Println(st.Head())

	st.Pop()
	fmt.Println(st.Head())

	fmt.Println("Hello World!")
	for i := 0; i < 100; i++ {
		fmt.Println("мяу")
	}
}
