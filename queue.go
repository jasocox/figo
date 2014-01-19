package queue

import (
	"container/list"
	"sync"
)

type Queue struct {
	l *list.List
}

type SQueue struct {
	q    *Queue
	lock sync.Mutex
}

func New() Queue {
	return Queue{list.New()}
}

func NewSync() SQueue {
	return SQueue{q: &Queue{list.New()}}
}

func (q Queue) Push(o interface{}) {
	q.l.PushBack(o)
}

func (q Queue) Pop() interface{} {
	e := q.l.Front()
	if e == nil {
		return nil
	}

	return q.l.Remove(e)
}

func (q Queue) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q Queue) Len() int {
	return q.l.Len()
}

func (q Queue) Front() *list.Element {
	return q.l.Front()
}

func (q Queue) Next(e *list.Element) *list.Element {
	if e == nil {
		return e
	}

	return e.Next()
}

func (q SQueue) Push(o interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.q.Push(o)
}

func (q SQueue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.q.Pop()
}

func (q SQueue) IsEmpty() bool {
	return q.q.IsEmpty()
}

func (q SQueue) Len() int {
	return q.q.Len()
}

func (q SQueue) Front() *list.Element {
	return q.q.Front()
}

func (q SQueue) Next(e *list.Element) *list.Element {
	return q.q.Next(e)
}
