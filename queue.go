// Package queue is a FIFO queue using the circularly linked list package in
// Go's standard library. There is also a thread safe version.
package queue

import (
	"container/list"
	"sync"
)

// Queue represents the FIFO queue.
type Queue struct {
	l *list.List
}

type AsyncQueue struct {
	q    *Queue
	lock sync.Mutex
}

// Returns an initialized Queue.
func New() Queue {
	return Queue{list.New()}
}

func NewSync() AsyncQueue {
	return AsyncQueue{q: &Queue{list.New()}}
}

// Pushes a new item to the back of the Queue.
func (q Queue) Push(o interface{}) {
	q.l.PushBack(o)
}

// Removes an item from the front of the Queue and returns it's value or nil.
func (q Queue) Pop() interface{} {
	e := q.l.Front()
	if e == nil {
		return nil
	}

	return q.l.Remove(e)
}

// Checks to see if the Queue is empty.
func (q Queue) IsEmpty() bool {
	return q.l.Len() == 0
}

// Returns the current length of the Queue.
func (q Queue) Len() int {
	return q.l.Len()
}

// Returns the item at the front of the Queue or nil.
// The item is a *list.Element from the 'container/list' package.
func (q Queue) Front() *list.Element {
	return q.l.Front()
}

// Returns the item after e or nil it is the last item or nil.
// The item is a *list.Element from the 'container/list' package.
// Even though it is possible to call e.Next() directly, don't. This behavior
// may not be supported moving forward.
func (q Queue) Next(e *list.Element) *list.Element {
	if e == nil {
		return e
	}

	return e.Next()
}

func (q AsyncQueue) Push(o interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q.Push(o)
}

func (q AsyncQueue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.q.Pop()
}

func (q AsyncQueue) IsEmpty() bool {
	return q.q.IsEmpty()
}

func (q AsyncQueue) Len() int {
	return q.q.Len()
}

func (q AsyncQueue) Front() *list.Element {
	return q.q.Front()
}

func (q AsyncQueue) Next(e *list.Element) *list.Element {
	return q.q.Next(e)
}
