package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)

	if q.length != 1 || q.head == nil || q.tail == nil || q.head.Value != 1 {
		t.Errorf("Enqueue failed for empty queue")
	}

	q.Enqueue(2)

	if q.length != 2 || q.head.Value != 1 || q.tail.Value != 2 {
		t.Errorf("Enqueue failed for enqueing to populated queue")
	}
}

func TestDequeue(t *testing.T) {
	q := Queue[int]{}

	q.Deque()

	if q.length != 0 || q.head != nil || q.tail != nil {
		t.Errorf("Dequeue failed to handle being called on empty queue")
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Deque()

	if q.length != 1 || q.head.Value != 2 || q.tail.Value != 2 {
		t.Errorf("Dequeue failed to handled dequeue of value")
	}
}
