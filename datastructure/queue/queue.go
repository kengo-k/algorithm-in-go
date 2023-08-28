package queue

import "container/list"

type LinkedListQueue struct {
	list *list.List
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		list: list.New(),
	}
}

func (q *LinkedListQueue) Enqueue(value any) {
	q.list.PushBack(value)
}

func (q *LinkedListQueue) Dequeue() any {
	element := q.list.Front()
	if element != nil {
		q.list.Remove(element)
		return element.Value
	}
	return nil
}

type ArrayQueue struct {
	older, younger []any
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

func (q *ArrayQueue) Enqueue(value any) {
	q.younger = append(q.younger, value)
}

func (q *ArrayQueue) Dequeue() any {
	if len(q.older) == 0 {
		if len(q.younger) == 0 {
			return nil
		}
		q.older, q.younger = q.younger, q.older
		reverse(q.older)
	}
	if len(q.older) == 0 {
		return nil
	}
	value := q.older[len(q.older)-1]
	q.older = q.older[:len(q.older)-1]
	return value
}

func reverse(a []any) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
