package queue

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Value T
	//Pointer to another node
	Next *Node[T]
}

type Queue[T any] struct {
	//Pointer to first and last node
	head   *Node[T]
	tail   *Node[T]
	length int
}

func (q *Queue[T]) Enqueue(item T) {
	//Take value, and return pointer to the memory where value is stored
	node := &Node[T]{Value: item}
	if q.length == 0 {
		q.head = node
		q.tail = q.head
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.length++
}

func (q *Queue[T]) Deque() {
	if q.length != 0 {
		q.head = q.head.Next
		q.length--
	}

	if q.head == nil {
		q.tail = nil
	}
}

func (q *Queue[T]) Peek() (T, error) {
	if q.head != nil {
		return q.head.Value, nil
	}

	var zero T
	return zero, errors.New("queue is empty")
}

func (q *Queue[T]) PrintQueue() {
	fmt.Printf("Head: %v\n", q.head)
	fmt.Printf("Tail: %v\n", q.tail)
	fmt.Printf("Length: %v\n", q.length)
	fmt.Println()
}

func (q *Queue[T]) String() string {
	var result string
	current := q.head
	for current != nil {
		result += fmt.Sprintf("%v -> ", current.Value)
		current = current.Next
	}
	result += "nil"
	return result
}
