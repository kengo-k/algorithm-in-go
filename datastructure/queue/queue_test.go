package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
	assert.Equal(t, q.Dequeue(), 3)
}

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, q.Dequeue(), 1)
	assert.Equal(t, q.Dequeue(), 2)
	assert.Equal(t, q.Dequeue(), 3)
}
